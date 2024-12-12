# stalewall-proxy
basic image proxy (with smartcrop support) for [stalewall](https://github.com/spacefall/stalewall-api) in pure go 

## Host
This repo is ready to be deployed to Vercel (serverless functions) and it is used for the main stalewall-proxy instance.
Alternatively, the repo includes a basic HTTP server (in go) that functions the same way as the function but is meant for development environments.

## Providers
Stalewall-proxy supports most of stalewall providers:
- Bing homepage image (`bing`)
- Windows Spotlight (`spotlight`)
- Chromecast 1 Ambient Mode (`chromecast`)
- Fire TV screensaver (`firetv`)
- Earth View by Google Earth (`earthview`)
- NASA Astronomy Picture of The Day (`apod`)

Note: value in parentheses is the one to use for prov query

Unsupported providers are the ones that aren't proxied by stalewall, like unsplash.

## Parameters
| Query | Value                             | Example    | Description                                                                                |
|-------|-----------------------------------|------------|--------------------------------------------------------------------------------------------|
| h     | (int>0)                           | ?h=1080    | Final height of proxied image                                                              |
| w     | (int>0)                           | ?w=1920    | Final width of proxied image                                                               |
| type  | specific value for each provider  | ?type=pp   | Tells the api the specific type of url to recompose                                        |
| prov  | provider of specified image       | ?prov=bing | Tells the api the image source to recompose the original url                               |
| id    | part of original url that changes | ?id=...    | Part of original url that identifies the image, it's encoded into base64 to prevent issues |
| q     | 0<=int<=100                       | ?q=92      | Sets image quality of proxied image (ignored if proxy is disabled)                         |
