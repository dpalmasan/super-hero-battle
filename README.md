# Super Hero Battles

## Assumptions

All the work can be tracked via the issues and commit messages. The only assumption I made is that
it is unlikely the heroes ids will be updated on a frequent cadence. Given this assumption, to generate
random ids, I just set a configuration, that by default will consider the ids in the [API official site](https://superheroapi.com/ids.html)

## Getting Started

For security purposes I didn't put any of my credentials (not `superheroapi` nor `mailgun`). I provided a `config.yaml.template` file, which must
be updated and renamed to `config.yaml`. The app was built in `Go 1.17`, I provided a `Dockerfile` so to easily run the project just run:

* `docker build -t super-hero-battle .`
* `docker run super-hero-battle`

If everything was setup correctly, some debug logs should appear and the output will be printed in `stdout`. If `Mailgun` was correctly configured, you
should get an email like:

![image](https://gist.githubusercontent.com/dpalmasan/103d61ae06cfd3e7dee7888b391c1792/raw/93794e3816d35829c12bac3dde1ff2f5c6f67198/battle-email.png)

If you'd like to run it locally with no docker, you will need to compile the sources:

* `CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/super-hero-battle ./cmd/super-hero-battle/`
* `./bin/super-hero-battle`

To run test suites: `ENV=test go test -v ./...`, output should look like:

```
=== RUN   TestGenerateRandomIds
--- PASS: TestGenerateRandomIds (0.00s)
PASS
ok      github.com/super-hero-battle/handlers   0.004s
=== RUN   TestBuildHeroTeam
--- PASS: TestBuildHeroTeam (0.00s)
=== RUN   TestUpdateFiliationCoefficient
--- PASS: TestUpdateFiliationCoefficient (0.00s)
PASS
ok 
```

## Development Process

All the development process can be found in the [issues](https://github.com/dpalmasan/super-hero-battle/issues?q=is%3Aissue).


## Some Caveats

I tried using SMTP method to send email via `mailgun` it worked locally, but it didn't work on Docker due to some certificates validation. I investigated a little bit, but it was faster just using the API.