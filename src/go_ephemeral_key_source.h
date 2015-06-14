#ifndef __GO_EPHEMERAL_KEY_SOURCE__H__
#define __GO_EPHEMERAL_KEY_SOURCE__H__

#include "net/quic/crypto/key_exchange.h"
#include "net/quic/crypto/ephemeral_key_source.h"
#include "net/quic/quic_time.h"
#include "base/memory/scoped_ptr.h"

class GoEphemeralKeySource : public net::EphemeralKeySource {
 public:
  GoEphemeralKeySource();

  virtual std::string CalculateForwardSecureKey(
      const net::KeyExchange* key_exchange,
      net::QuicRandom* rand,
      net::QuicTime now,
      base::StringPiece peer_public_value,
      std::string* public_value) override;

 private:
  scoped_ptr<net::KeyExchange> forward_secure_key_exchange_;
  net::QuicTime key_created_time_;
};

#endif  // __GO_EPHEMERAL_KEY_SOURCE__H__
