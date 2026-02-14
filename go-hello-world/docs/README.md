# Project Documentation

This repository contains a full-stack web application built with Go and Gin. The app uses server-side rendering (SSR) with Go HTML templates and Bootstrap for styling. The project is designed to run on Linux and will be deployed to an AWS EC2 instance later in the semester.

---

## Technology Stack

### Backend technology stack
- Backend Language: Go
- Backend Framework: Gin (HTTP routing + middleware)
- Runtime / Server: Go net/http server via Gin
- Database: Not implemented yet
- Deployment Target: AWS EC2 (Linux)

### Frontend technology stack
- Templating: Go html/template (rendered through Gin)
- UX/UI Framework: Bootstrap (via CDN in templates)

### Testing Frameworks
- Unit / Integration Testing: Go built-in testing package
- HTTP Route Testing: net/http/httptest
- Test Command: `go test ./...`

### Manual Deployment
- Application runs as a compiled Go binary.
- Will be deployed to AWS EC2 (Linux) later in the semester.
- Reverse proxy (e.g., nginx) may be added in production.

---

## Team Workflow

- Development Mode: Solo developer
- Repository Strategy: Single repository (no forks / pull requests)
- Branching: Work is done on main for now with the ability to switch to feature branches
- Commits: Small and frequent commits with clear messages 
