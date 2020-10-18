1. Change input to NextTag to be a reader; give tag an internal byte buffer -> This the whole point.
2. Decide if we're sticking keep the raw input tag attached to the tag, or if we should keep the extra fields and re-serialize
3. Allow more configuration options from the CLI - especially autoformat
