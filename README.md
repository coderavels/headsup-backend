# headsup-backend

## Endpoints
/decks
- GET - get a list of deck

/decks/:deck/cards
- GET - get list of cards for a deck

## Deploy changes in handlers to lambdas
1. `make build && make build-lambda && sam deploy`

## Run the lambda locally
1. `make build && make build-lambda`
2. sam local start-api

## Makefile
1. ensure the `BINARY_NAME` in `Makefile` matches the `Resources.<name>.Properties.Handler` in `template.yml`
2. ensure the `ZIP_NAME` is `Makefile` matches the zip file name in `Resources.<name>.Properties.CodeUri` in `template.yml`
### make commands
1. `make build`: build the binary
2. `make build-lambda`: build the deployable zip for lambda deployment

## Set up aws sam for deployment from local
1. follow first party installation steps (without homebrew) as mentioned [here](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html)
2. ```
    ❯ which sam
    /usr/local/bin/sam
    ❯ sam --version
    SAM CLI, version 1.97.0
    ```

## Set up sso session with AWS cli (first update aws cli)
1. `aws configure sso`
2. 
```
❯ aws configure sso
SSO session name (Recommended): <session-name>
SSO start URL [None]: https://example.awsapps.com/start#
SSO region [None]: us-east-1
SSO registration scopes [sso:account:access]: sso:account:access
Attempting to automatically open the SSO authorization page in your default browser.
If the browser does not open or you wish to use a different device to authorize this request, open the following URL:

https://device.sso.us-east-1.amazonaws.com/

Then enter the code:

TPBC-TGBM
The only AWS account available to you is: 
Using the account ID 
The only role available to you is: 
Using the role name ""
CLI default client Region [us-west-2]:
CLI default output format [None]: json
CLI profile name [AdministratorAccess]: <profile-name>

To use this profile, specify the profile name using --profile, as shown:

aws s3 ls --profile <profile-name>
```
3. `aws sso login --sso-session <profile-name>`
