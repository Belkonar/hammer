# Hammer

Hammer is an API gateway with a focus on transparent authorization.

It uses prefix paths to work. It's crazy annoying but needed to properly handle
APIs.

It's designed as two separate services. If this takes off I'll end up using
Cobra, but for now it'll be fairly simple as two executables.

## Database

Fairly simply, the database is just etcd. That way we can easily update the
gateway in real time.

## Authorization

This will use FGA/OpenFGA to handle the authorization layer. This may become
more flexible over time, but for now that'll be it.

## UI?

Not currently. API only.

## Why does this even exist?

Mostly because I wanted an API gateway that wasn't a complete pain to use, and
that also supported real time updates of configuration. Existing solutions
didn't have great ability for me to hook in the way I need to, or are "fremium"
solutions that are useless without paying a ton of money. This project will be
always be fully free (as in free beer) and open source.
