# Following REST API tutorial - Flash cards API

## Diffs:

- Use Gin
- Use Zap instead of Logrus
  - How to substitue Gin inner logger? Looks like it is harcoded into Gin Engine.
- Entities:
  - User
  - Deck
  - Card
- Auth - Basic

## Find more about

- Project structure: https://github.com/golang-standards/project-layout
- Testing (testify / Mocks / Table Tests)
- Makefile tool
- golang-migrate tool

## Optional

- Implement Flash cards "shuffler"
- Move Shuffler to another service
- Write Protobufs
- Connect to "shuffler" via gRPC
