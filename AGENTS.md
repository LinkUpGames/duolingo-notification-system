## Monorepo Agent Guidelines

This document provides instructions for AI agents working in this repository.

### Build, Lint, and Test

- **Client (TypeScript/React):**
  - `npm run build`: Compile TypeScript.
  - `npm test`: Run Prettier, XO (lint), and Ava (tests).
  - To run a single test: `npx ava client/source/test.tsx` (replace with your file).
- **Scorer (Python/C):**
  - `cd scorer/extensions/algorithm && make all`: Build and install the C extension.
  - `pip install -r scorer/requirements.txt`: Install Python dependencies.
  - `python scorer/src/test.py`: Run scorer tests.
- **Server (Go):**
  - `go build`: Compile the server.
  - `go test ./...`: Run all Go tests.

### Code Style

- **General:** Follow existing code style. Use Prettier for formatting.
- **TypeScript/React:** Adhere to `xo-react` rules. Use functional components with hooks.
- **Go:** Follow standard Go conventions (gofmt).
- **Python:** Follow PEP 8.
- **Error Handling:** Handle errors explicitly; no silent failures.
- **Dependencies:** Add new dependencies to the appropriate file (`package.json`, `requirements.txt`, `go.mod`).

