# Commons Module
Contains the protobuf stuff that determines the gRPC communication
and some functionalities that are used by both client and service.

## Protobuf compilation
From the root folder of this project (parent folder of this module),
execute the following command:
```sh
protoc --go_out=commons commons/proto/*.proto --go-grpc_out=commons
```
This will generate a folder "pb" with the golang protobuf files, check if their
package name is: "block". If is different from that change it to "block".

## Package/structure summary
- models:
  - Definition of the block struct
  - Serialization and deserialization functions
- pb:
  - Generated files by protoc
- proto:
  - Protobuf definition
- testUtils:
  - Common unity test functions
- blocksFlags:
  - Flags/constants which defines routes, domains, names...