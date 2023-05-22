# demoScrape2

The stats collector library for CSC demo files.

## Deploy Checklist

After making changes to this repo, follow the following steps:

- [ ] Verify all new JSON field names are in camelcase and are public/exported (this matters later)
- [ ] Create release with a new tag using [semantic versioning](https://semver.org/), also use this page to also auto-generate release notes

## Post Deploy

- Follow the instructions [here](https://github.com/csconfederation/csgo-demo-worker/blob/master/README.md#deploy-checklist) to push the latest parser changes to the demo worker
- Follow the instructions [here](https://github.com/csconfederation/CSC-Stats#add-migrations) to update the prisma schema & auto generate GraphQL queries in the stats-api
