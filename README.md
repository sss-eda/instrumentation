# Instrumentation
This repository represents a domain of the business. This domain is the **Instrumentation** domain and contains a multitude of bounded contexts.

## Entities:
* Instrument - Entity within Site Aggregate
* Site - Aggregate Root
* InstrumetType - A separate Aggregate, gets referenced by Site

## Usecases
1. Add/Remove an Instrument to/from a Site.
2. Register a new InstrumentType.
3. Add/Remove a Site.
4. Move an instrument from one site to another.
5. Change an instruments type...? -> Impossible. The instrument's type is inherent to its identity.

## Aggregates
Contains the following aggregates:
* Instrument
    - VO: Name
    - VO: Iteration
* Site
    - VO: Name
    - VO: Abbreviation
* InstrumentType
    - VO: Name
    - VO: Abbreviation