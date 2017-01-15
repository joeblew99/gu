Getting Started
===============

Gu is fundamentally a library built to handle rendering of html on either front or backend.

When creating Gu, my main focus was on creating a solution that did not bind itself tightly to the perculiarities of either the browser or server, but allow each structures to define for themselves how they should be rendered and also allow developers to define how those components and other content should be rendered, and ensuring it works both on the server and browser, which will be often referred to as the backend and frontend.

The concepts in Gu are practially simple and rely majorly on a functional and interface based constructs, where structures can define the markup to be rendered by implementing certain interfaces and also elivate themselves as reactive by implement others. These approach provides a high level of simplicity and ease in both thinking and use without enforcing any rigid rules.

*Gu is in no way a Flux-like framework. It is just a library. It does provide complex structures and layed down paths by which such can be attained, it simply provides a baseline to render the desire output and gives the freedom for the developer to determine how their application data flow should works.*

The Guide
---------
To ensure this getting started is as simple as possible, we will demonstrate in code
how to build a simple component and how that component can be deployed for use in a page.


Last Note
---------

Please ensure to check out the [Examples](../examples) directory of the source to both study and understand the code for the examples, and accompany it with these explanations to help provide insightful understanding of how `Gu` works.

Please feel free to issue PRs and file issues on suggestions, changes and improvement or just questions to help you better understand Gu.

God bless.
