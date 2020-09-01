# Youtube-Channel-Badge

A shields.io API for your youtube channel to protect your api key.

## Table of contents

- [Why](https://github.com/ntec-io/youtube-channel-badge#Why)
- [Features](https://github.com/ntec-io/youtube-channel-badge#Features)
- [Examples](https://github.com/ntec-io/youtube-channel-badge#Examples)
  - [URLs](https://github.com/ntec-io/youtube-channel-badge#URLs)
  - [Styles](https://github.com/ntec-io/youtube-channel-badge#Styles)
- [Configuration](https://github.com/ntec-io/youtube-channel-badge#Configuration)
  - Getting a google api key & activate youtube api
  - Find your youtube channel id
- [Deployment](https://github.com/ntec-io/youtube-channel-badge#Deployment)
  - Step by Step
- [Usage](https://github.com/ntec-io/youtube-channel-badge#Usage)

## Why

You can't publicly access the google/youtube api without an api key. That's why there is no shields.io badge for youtube channels. This repo aims to be the solution to this, because you deploy your own public api which only fetches your channel.

## Features

- Cache time to prevent rate limiting
- Secure api key due env variable

## Examples

### Urls

| Name | Example | Description | URL |
| --- | --- | --- | --- |
| Subscriber count | ![Custom badge](https://img.shields.io/endpoint?url=https%3A%2F%2Fyoutube-channel-badge.ngoldack.vercel.app%2Fapi%2Fsubscriber) | Shows the amount of subscribers | /subscriber |
| View count | ![Custom badge](https://img.shields.io/endpoint?url=https%3A%2F%2Fyoutube-channel-badge.ngoldack.vercel.app%2Fapi%2Fviews) | Shows the amount of views of all videos of the channel | /views |
| Video count | ![Custom badge](https://img.shields.io/endpoint?url=https%3A%2F%2Fyoutube-channel-badge.ngoldack.vercel.app%2Fapi%2Fvideos) | Shows the amount of videos on the channel | /videos |
| Comment count | ![Custom badge](https://img.shields.io/endpoint?url=https%3A%2F%2Fyoutube-channel-badge.ngoldack.vercel.app%2Fapi%2Fcomments) | Shows the amount of youtube comments of the channel | /comments |

### Styles

| Name | Examples |
| --- | --- |
| plastic | ![Custom badge](https://img.shields.io/endpoint?style=plastic&url=https%3A%2F%2Fyoutube-channel-badge.ngoldack.vercel.app%2Fapi%2Fsubscriber) |
| flat | ![Custom badge](https://img.shields.io/endpoint?url=https%3A%2F%2Fyoutube-channel-badge.ngoldack.vercel.app%2Fapi%2Fsubscriber) |
| flat-square | ![Custom badge](https://img.shields.io/endpoint?style=flat-square&url=https%3A%2F%2Fyoutube-channel-badge.ngoldack.vercel.app%2Fapi%2Fsubscriber) |
| for-the-badge | ![Custom badge](https://img.shields.io/endpoint?style=for-the-badge&url=https%3A%2F%2Fyoutube-channel-badge.ngoldack.vercel.app%2Fapi%2Fsubscriber) |
| social | ![Custom badge](https://img.shields.io/endpoint?style=social&url=https%3A%2F%2Fyoutube-channel-badge.ngoldack.vercel.app%2Fapi%2Fsubscriber) |

All examples are live and based on my youtube channel: [ntec.io](https://www.youtube.com/channel/UCWlzQ62V0vhtH1XkgN4hS2A)

## Configuration

| Variable Name | Description | Required |
| --- | --- | --- |
| API_KEY | Your google api key. See section below on how to get one | **TRUE**
| CHANNEL_ID | The youtube channel id of which the stats should be processed | **TRUE** |
| CACHE_TIME | Time in seconds on how long the last result is cached. Default is 300. | **FALSE** |

### Getting a google api key & activate youtube api

1. Go to [google developer console](https://console.developers.google.com/)

2. Log In to your google account

3. Click on **Select a project** (Looks like a dropdown button)

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/apikey01.png?raw=true)

4. Create a new project by clicking on the **New Project** button

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/apikey02.png?raw=true)

5. Give your project a name (or leave it to the generated one) and create your project by pressing on **CREATE**

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/apikey03.png?raw=true)

6. On the next window, click on **+ ENABLE APIS AND SERVICES** in the top

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/apikey04.png?raw=true)

7. Search for **Youtube** in the search bar and click on **YoTube Data API v3**

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/apikey05.png?raw=true)

8. Enable the API by clicking on **ENABLE**

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/apikey06.png?raw=true)

9. Click on the 3 lines in the top left corner to open the side menu and hover over **APIs & Services** and then click on **Credentials**

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/apikey07.png?raw=true)

10. Click on **CREATE CREDENTIALS** and then on **API key**

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/apikey08.png?raw=true)

11. Copy your API key

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/apikey09.png?raw=true)

### Find your youtube channel id

1. Goto your [advanced account settings](https://www.youtube.com/account_advanced)
2. Copy your channel id

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/channelid.png?raw=true)

## Deployment
It is neccessary to deploy this repo by yourself to get a domain for the desired youtube channel.
An easy and free way to deploy this repo is by using vercel. Just use the button below.

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/import/git?s=https%3A%2F%2Fgithub.com%2Fntec-io%2Fyoutube-channel-badge&env=API_KEY,CHANNEL_ID,CACHE_TIME&envDescription=Find%20information%20on%20how%20to%20get%20these%20in%20the%20readme&envLink=https%3A%2F%2Fgithub.com%2Fntec-io%2Fyoutube-channel-badge%23configuration)

### Step by step

1. Click on the **deploy** button above or click [here](https://vercel.com/import/git?s=https%3A%2F%2Fgithub.com%2Fntec-io%2Fyoutube-channel-badge&env=API_KEY,CHANNEL_ID,CACHE_TIME&envDescription=Find%20information%20on%20how%20to%20get%20these%20in%20the%20readme&envLink=https%3A%2F%2Fgithub.com%2Fntec-io%2Fyoutube-channel-badge%23configuration)

2. Click on **Continue**

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/deploy01.png?raw=true)

3. Click on **no**

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/deploy02.png?raw=true)

4. Select your user account and click on **continue**

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/deploy03.png?raw=true)

5. Fork the repo with a click on **continue**

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/deploy04.png?raw=true)

6. Click again on **Continue**

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/deploy05.png?raw=true)

7. Enter your API_KEY, CHANNEL_ID and CACHE_TIME. Check [Configuration](https://github.com/ntec-io/youtube-channel-badge#Configuration) if you need help getting those things.

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/deploy06.png?raw=true)

8. Click on **Deploy** and your deployment should finish after a short time

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/deploy07.png?raw=true)

## Usage

1. Deploy the repo and get your url from **domains**

    Example in vercel:

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/deployUrl.png?raw=true)

2. Go to shields.io/endpoint and scroll down

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/usage01.png?raw=true)

3. Add your vercel url to the **url** field
4. Add your wanted url suffix:
    - to get the subscriber count: add api/subscriber

    - to get the view count: add api/views

    - to get the comment count: add api/comments

    - to get the video count: add api/videos

    It should look like this:

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/usage02.png?raw=true)

5. Override the labels/colors/logo if you want (optional)
6. Click on the button to copy the badge url

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/usage03.png?raw=true)

7. Enjoy! :)
