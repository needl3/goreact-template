# GoReact
Template for building websites with Golang and React with Hexagonal Architecture.

## Features
1. Full on backend with stdlib server and router. No external routing framework required.
2. Ready to use user authentication with GoogleOAuth
3. Ready to use database with PostgreSQL. Uses (pgxpool)[https://github.com/jackc/pgx] for database driver 
4. Use raw queries to maintain database. Including migration and rollback.
5. Ready to use frontend with (React)[https://reactjs.org/]❤️(Typescript)[https://www.typescriptlang.org/] (ESBuild)[https://esbuild.github.io/] and (TailwindCSS)[https://tailwindcss.com/]
6. Check session and allow UI routes from go side.

## Stack
1. (Golang)[https://golang.org/doc/install]
2. (PostgreSQL)[https://www.postgresql.org/]
3. (React)[https://reactjs.org/]
4. (ESBuild)[https://esbuild.github.io/]
5. (TailwindCSS)[https://tailwindcss.com/docs/installation]
6. (Typescript)[https://www.typescriptlang.org/]

## How to use
1. Clone (this)[https://github.com/needl3/goreact-template] repository
2. Populate env file
3. Make sure you have database mentioned on env file
4. Make sure you have (yarn)[https://yarnpkg.com/] installed
5. Run `make init`. It will initialize backend dependencies and frontend using yarn
6. Migrate database with `make migrate`
7. Run backend with `make backend-dev`
7. Run frontend with `make frontend-dev`
