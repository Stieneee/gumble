package gumble

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/stieneee/gumble/gumble/proto/MumbleUDPProto"
	"google.golang.org/protobuf/proto"
)

// PingResponse contains information about a server that responded to a UDP
// ping packet.
type PingResponse struct {
	// The address of the pinged server.
	Address *net.UDPAddr
	// The round-trip time from the client to the server.
	Ping time.Duration
	// The server's version. Only the Version field and SemanticVersion method of
	// the value will be valid.
	Version Version
	// The number users currently connected to the server.
	ConnectedUsers int
	// The maximum number of users that can connect to the server.
	MaximumUsers int
	// The maximum audio bitrate per user for the server.
	MaximumBitrate int
}

// Ping sends a UDP ping packet to the given server. If interval is positive,
// the packet is retransmitted at every interval.
//
// This function will fail to return a proper response for servers using newer
// versions of the Mumble UDP protocol. Use PingProto instead.
//
// Returns a PingResponse and nil on success. The function will return nil and
// an error if a valid response is not received after the given timeout.
func Ping(address string, interval, timeout time.Duration) (*PingResponse, error) {
	if timeout < 0 {
		return nil, errors.New("gumble: timeout must be positive")
	}
	deadline := time.Now().Add(timeout)
	conn, err := net.DialTimeout("udp", address, timeout)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	conn.SetReadDeadline(deadline)

	var (
		idsLock sync.Mutex
		ids     = make(map[string]time.Time)
	)

	buildSendPacket := func() {
		var packet [12]byte
		if _, err := rand.Read(packet[4:]); err != nil {
			return
		}
		id := string(packet[4:])
		idsLock.Lock()
		ids[id] = time.Now()
		idsLock.Unlock()
		conn.Write(packet[:])
	}

	if interval > 0 {
		end := make(chan struct{})
		defer close(end)
		go func() {
			ticker := time.NewTicker(interval)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					buildSendPacket()
				case <-end:
					return
				}
			}
		}()
	}

	buildSendPacket()

	for {
		var incoming [24]byte
		if _, err := io.ReadFull(conn, incoming[:]); err != nil {
			return nil, err
		}
		id := string(incoming[4:12])
		idsLock.Lock()
		sendTime, ok := ids[id]
		idsLock.Unlock()
		if !ok {
			continue
		}
		v := binary.BigEndian.Uint32(incoming[0:])
		major := uint16(v>>16) & 0xFFFF
		minor := uint16(v>>8) & 0xFF
		patch := uint16(v) & 0xFF

		return &PingResponse{
			Address: conn.RemoteAddr().(*net.UDPAddr),
			Ping:    time.Since(sendTime),
			Version: Version{
				Version: (uint64(major) << 48) | (uint64(minor) << 32) | (uint64(patch) << 16),
			},
			ConnectedUsers: int(binary.BigEndian.Uint32(incoming[12:])),
			MaximumUsers:   int(binary.BigEndian.Uint32(incoming[16:])),
			MaximumBitrate: int(binary.BigEndian.Uint32(incoming[20:])),
		}, nil
	}
}

// Ping sends a UDP ping packet to the given server. If interval is positive,
// the packet is retransmitted at every interval. This function uses the Mumble
// UDP protobuf protocol.
//
// Returns a PingResponse and nil on success. The function will return nil and
// an error if a valid response is not received after the given timeout.
func PingProto(address string, interval, timeout time.Duration) (*PingResponse, error) {
	if timeout < 0 {
		return nil, errors.New("gumble: timeout must be positive")
	}
	deadline := time.Now().Add(timeout)
	conn, err := net.DialTimeout("udp", address, timeout)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	conn.SetReadDeadline(deadline)

	var (
		idsLock sync.Mutex
		ids     = make(map[string]time.Time)
	)

	buildSendPacket := func() {
		p := MumbleUDPProto.Ping{
			Timestamp:                  uint64(time.Now().Unix()),
			RequestExtendedInformation: true,
		}

		packet, err := proto.Marshal(&p)
		if err != nil {
			log.Println("Failed to marshal ping packet")
			return
		}

		var header [1]byte
		header[0] = 0x01
		packet = append(header[:], packet...)

		id := fmt.Sprint(p.GetTimestamp())
		idsLock.Lock()
		ids[id] = time.Now()
		idsLock.Unlock()

		conn.Write(packet[:])
	}

	if interval > 0 {
		end := make(chan struct{})
		defer close(end)
		go func() {
			ticker := time.NewTicker(interval)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					buildSendPacket()
				case <-end:
					return
				}
			}
		}()
	}

	buildSendPacket()

	for {
		var incoming []byte
		buf := make([]byte, 24) // Adjust the buffer size as needed
		for {
			n, err := conn.Read(buf)
			if err != nil {
				return nil, err
			}
			incoming = append(incoming, buf[:n]...)
			if n < len(buf) {
				break
			}
		}

		p := MumbleUDPProto.Ping{}
		err := proto.Unmarshal(incoming[1:], &p)
		if err != nil {
			log.Println("Failed to unmarshal ping packet")
			continue
		}
		// Unmarshal succeeded, process the packet
		id := fmt.Sprint(p.GetTimestamp())
		idsLock.Lock()
		sendTime, ok := ids[id]
		idsLock.Unlock()
		if !ok {
			log.Println("Received unknown ping response")
			continue
		}

		return &PingResponse{
			Address: conn.RemoteAddr().(*net.UDPAddr),
			Ping:    time.Since(sendTime),
			Version: Version{
				Version: p.GetServerVersionV2(),
			},
			ConnectedUsers: int(p.UserCount),
			MaximumUsers:   int(p.MaxUserCount),
			MaximumBitrate: int(p.MaxBandwidthPerUser),
		}, nil
	}
}
