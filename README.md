<p align="center">
  <a href="https://github.com/fastenhealth/fasten-onprem">
  <img width="400" alt="fasten_view" src="frontend/src/assets/banner/banner.png">
  </a>
</p>

# Fasten - On Premise/Self-Hosted

[![CI](https://github.com/fastenhealth/fasten-onprem/actions/workflows/ci.yaml/badge.svg)](https://github.com/fastenhealth/fasten-onprem/actions/workflows/ci.yaml)
[![GitHub license](https://img.shields.io/github/license/fastenhealth/fasten-onprem?style=flat-square)](https://github.com/fastenhealth/fasten-onprem/blob/main/LICENSE.md)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/fastenhealth/fasten-onprem?style=flat-square)](https://github.com/fastenhealth/fasten-onprem/releases/latest)
[![Discord Join](https://img.shields.io/badge/discord-join-blueviolet?style=flat-square&logo=discord)](https://discord.gg/Bykz6BAN8p)

**Fasten securely connects your healthcare providers together, creating a personal health record that never leaves your hands**

> NOTE: Fasten is a Work-in-Progress and cannot communicate with your actual healthcare providers (yet).
> Please fill out this [Google Form](https://forms.gle/SNsYX9BNMXB6TuTw6) if you'd like to be kept up-to-date on Fasten 

# Introduction

Like many of you, I've worked for many companies over my career. In that time, I've had multiple health, vision and dental 
insurance providers, and visited many different clinics, hospitals and labs to get procedures & tests done.

Recently I had a semi-serious medical issue, and I realized that my medical history (and the medical history of my family members) 
is a lot more complicated than I realized and distributed across the many healthcare providers I've used over the years. 
I wanted a single (private) location to store our medical records, and I just couldn't find any software that worked as I'd like:

- self-hosted/offline - this is my medical history, I'm not willing to give it to some random multi-national corporation to data-mine and sell 
- It should aggregate my data from multiple healthcare providers (insurance companies, hospital networks, clinics, labs) across multiple industries (vision, dental, medical) -- all in one dashboard 
- automatic - it should pull my EMR (electronic medical record) directly from my insurance provider/clinic/hospital network - I dont want to scan/OCR physical documents (unless I have to)
- open source - the code should be available for contributions & auditing

So, I built it

**Fasten is an open-source, self-hosted, personal/family electronic medical record aggregator, designed to integrate with 1000's of insurances/hospitals/clinics**

Here's a couple of screenshots that'll give you an idea of what it looks like:


<p align="center">
  <a href="https://imgur.com/a/vfgojBD">
  <img alt="fasten_view" src="https://i.imgur.com/UaZyEbN.png">
  </a>
  <br/>
  <a href="https://imgur.com/a/vfgojBD">See more Fasten screenshots</a>
</p>

# Features

It's pretty basic right now, but it's designed with a easily extensible core around a solid foundation:

- Self-hosted
- Designed for families, not Clinics (unlike OpenEMR and other popular EMR systems)
- Supports the Medical industry's (semi-standard) FHIR protocol
- Uses OAuth2 (Smart-on-FHIR) authentication (no passwords necessary)
- Uses OAuth's `offline_access` scope (where possible) to automatically pull changes/updates
- Multi-user support for household/family use
- (Future) Dashboards & tracking for diagnostic tests
- (Future) Integration with smart-devices & wearables

# Getting Started

## Instructions

Run the following commands to download and start the Fasten docker container.
```
docker pull ghcr.io/fastenhealth/fasten-onprem:main 
docker run --rm -p 9090:8080 ghcr.io/fastenhealth/fasten-onprem:main 
```

Next, open a browser to `http://localhost:9090`

At this point you'll be redirected to the login page.

### Logging In

Before you can use the Fasten BETA, you'll need to [Create an Account](http://localhost:9090/web/auth/signup).

It can be as simple as
- **Username:** `testuser`
- **Password:** `testuser`


## Usage


# FAQ's

See [FAQs](https://github.com/fastenhealth/docs/blob/main/FAQs.md) for common questions (& answers) regarding Fasten

# Contributing

Please see the [CONTRIBUTING.md](CONTRIBUTING.md) for instructions for how to develop and contribute to the Fasten codebase.

Work your magic and then submit a pull request. We love pull requests!

If you find the documentation lacking, help us out and update this README.md. If you don't have the time to work on Fasten, but found something we should know about, please submit an issue.

# Versioning

We use SemVer for versioning. For the versions available, see the tags on this repository.

# Authors

Jason Kulatunga - Initial Development - @AnalogJ

# Licenses

- GPLv3

# Sponsors

Fasten is only possible with the help of my [Github Sponsors](https://github.com/sponsors/AnalogJ/) and Reddit supporters!
