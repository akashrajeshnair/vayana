This is a full-stack book reading tracker app (like Goodreads).
- Backend: Spring Boot, PostgreSQL, JWT auth
- Frontend: NextJS, React, TypeScript, Tailwind CSS
- AWS Lambda: welcome email on signup, weekly reading summary email (via AWS SES)
- Code should be simple, readable, and written at a junior engineer level
- Avoid over-engineering. Prefer clarity over cleverness.
- Backend follows a standard Controller -> Service -> Repository pattern
- No Lombok. Write out getters/setters explicitly.
- REST API returns standard JSON responses

The following points are only relevant at time of deployment, so don't worry about them during development:
- Backend is deployed as a plain JAR on EC2, managed by systemd
- CI/CD uses GitHub Actions
- Frontend is deployed on AWS Amplify
- No custom domain, use raw IPs/URLs

The project structure is as follows:
vayana/
├── backend/        (Spring Boot)
├── frontend/       (React + TypeScript)
└── lambda/         (two separate Lambda functions)