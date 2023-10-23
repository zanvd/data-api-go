<p align="center">
    <a href="https://gatehub.net">
      <img src="https://cdn.gatehub.net/img/gatehub_logo_blue.svg" alt="GateHub"/ width="500px">
    </a>
</p>
<h3 align="center">GateHub Data API</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]() [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

</div>

---

<p align="center"> The GateHub Data API v2 provides access to information about changes in the XRP Ledger, including transaction history and processed analytical data. This information is stored in a dedicated database, which frees rippled servers to keep fewer historical ledger versions.
    <br> 
GateHub provides a live instance of the Data API with as complete a transaction record as possible at the following address: <br /><br />
<table>
  <thead>
    <tr>
      <th colspan="2">Endpoints</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>XRPL Livenet</td>
      <td><a href="https://data.gatehub.net" target="_blank">https://data.gatehub.net</a></td>
    </tr>
    <tr>
      <td>XRPL Testnet</td>
      <td><a href="https://data.sandbox.gatehub.net" target="_blank">https://data.sandbox.gatehub.net</a></td>
    </tr>
  </tbody>
</table>


## 📝 Table of Contents

- [Problem Statement](#problem_statement)
- [Idea / Solution](#idea)
- [Dependencies / Limitations](#limitations)
- [Setting up a local environment](#getting_started)
- [Usage](#usage)
- [Technology Stack](#tech_stack)
- [Contributing](../CONTRIBUTING.md)
- [Authors](#authors)

## 🧐 Problem Statement <a name = "problem_statement"></a>

### Background
The Ripple Data API has long served as a crucial bridge for developers and businesses to interact with and retrieve essential XRP Ledger data. This API provided a streamlined way to access ledger, transaction, and account data without the overhead of running a full node. However, with its deprecation and eventual shutdown, a significant gap has emerged. Developers, businesses, and enthusiasts who depended on this API are now left without an easy-to-use interface to fetch and analyze XRP Ledger data.

### Desired State
The ideal solution would be a robust, scalable, and maintainable API that mirrors the functionality provided by the now-deprecated Ripple Data API. This replacement should:

1. **Ease of Use**: Be intuitive and accessible, even for those new to the XRP Ledger.
2. **Comprehensive Data Access**: Offer comprehensive access to ledger, transaction, and account data, ensuring users don't feel a functionality gap from the original Ripple Data API.
3. **Scalability**: Handle a large number of requests efficiently, ensuring fast response times and high availability.
4. **Open-Source and Community-Driven**: Allow the community to contribute, ensuring the API evolves in response to user needs and stays up-to-date with any changes or advancements in the XRP Ledger.
5. **Transparency**: Offer detailed documentation and usage examples, promoting transparency and assisting developers in integrating with their applications.

Once this solution is implemented, developers and businesses will be empowered to seamlessly integrate XRP Ledger data into their applications and services, fostering innovation and growth within the XRP ecosystem.

## 💡 Idea / Solution <a name = "idea"></a>

In light of the deprecation of the Ripple Data API, the XRP community faces the challenge of finding a suitable, modern, and scalable alternative. Our solution, the XRP Ledger Data Service, is built from the ground up, leveraging cutting-edge technologies and architecture to ensure reliability, scalability, and maintainability.

### Key Features
1. **Powered by BigQuery**: Our backend relies on Google BigQuery, ensuring high-speed data retrieval, powerful analytics, and a scalable database solution.
2. **Go and Protocol Buffers**: By leveraging the Go programming language, known for its performance and efficiency, along with Protocol Buffers for serialization, we ensure a lightweight, high-performance backend.
3. **Dual API Access**:
    * **gRPC Access**: For clients seeking high performance and efficiency, our service offers gRPC endpoints.
    * **RESTful Access**: For broad compatibility and easy integration, we also expose a RESTful API, mirroring the familiar interface of the original Ripple Data API.
4. **Microservice Architecture**: Our system is designed as a set of microservices, ensuring scalability, maintainability, and resilience. Each service is purpose-built, ensuring optimal performance for its designated task.
5. **Kubernetes and CI/CD**: Deployed within a Kubernetes cluster, our service is designed for high availability and automatic scaling. Continuous Integration and Continuous Deployment (CI/CD) pipelines ensure that updates are seamless, with zero downtime.
6. **Security First**: We prioritize the security of our service. All containers undergo rigorous security scans. Additionally, the entire ecosystem is continuously monitored for vulnerabilities, ensuring the protection of user data and service integrity.

### Benefits
* **Scalability**: Built for the modern web, our service can handle the demands of thousands of simultaneous users, scaling as needed.
* **Reliability**: Leveraging Kubernetes, we ensure that our service is always available, automatically healing and scaling as required.
* **Future-Proof**: With an open-source foundation and community-driven focus, our solution will continually evolve, staying current with the latest advancements in the XRP Ledger and the broader tech landscape.
* **Easy Transition**: For those familiar with the Ripple Data API, transitioning to our service will be smooth, with familiar endpoints and enhanced performance.

The GateHub Data API isn't just a replacement; it's an upgrade. Whether you're a developer, business, or enthusiast, our service offers a next-generation solution for all your XRP Ledger data needs.

This is a broad-strokes overview, and you might want to delve deeper into specifics, especially if there are unique features or innovations you're introducing. Nonetheless, we hope this provides a strong foundation for your project's documentation!


## ⛓️ Dependencies / Limitations <a name = "limitations"></a>
1. **Rate Limiting:** All users will need to provide an API key to access the service. For users on the free tier, there's a limit of 100 requests per day with the provided API key. To increase this limit or for more comprehensive access plans, users should consult our pricing and subscription details available on the GateHub page.
2. **Dependency on BigQuery:** While Google BigQuery offers excellent performance and scalability, any potential outages or disruptions to BigQuery could impact our service's availability.
3. **Data Freshness:** There might be a slight delay between real-time XRP Ledger events and their appearance in the API due to processing and synchronization times.
4. **Maintenance Windows:** We strive for maximum uptime, but there might be occasional maintenance windows leading to temporary unavailability. Users will be notified in advance of scheduled maintenance. For real-time updates and to stay informed about any service disruptions or maintenance events, users can subscribe to the GateHub status page, accessible at status.gatehub.net.
5. **gRPC/REST Compatibility:** Although we offer both gRPC and RESTful endpoints, there could be subtle differences in the response or behavior due to the inherent differences in the protocols.
6. **Feature Parity:** While our goal is to replicate the functionalities of the deprecated Ripple Data API, some lesser-used features might be introduced later based on user demand.
7. **Resource Intensive Queries:** Complex or broad queries might face longer processing times or might be truncated to ensure system stability.
8. **Documentation Discrepancies:** As our service continues to evolve and improve, there might be occasional discrepancies between the live features and the documentation. We are committed to keeping the documentation up-to-date to reflect the current state of the service.


## 🏁 Getting Started <a name = "getting_started"></a>
**TBD**

### Prerequisites
**TBD**

### Installing
**TBD**

## 🎈 Documentation <a name="usage"></a>

Check GateHub official documentation at https://docs.gatehub.net

## ⛏️ Built With <a name = "tech_stack"></a>

- [BigQuery](https://cloud.google.com/bigquery) - Database
- [Protocol Buffers](https://protobuf.dev) - Data Definition
- [Go](https://go.dev) - Programing Language
- [gRPC Gateway](https://go.dev) - gRPC to Rest API

## ✍️ Authors <a name = "authors"></a>

- [@tadejgolobic](https://github.com/tadejgolobic) - Idea & Initial work
- [@rokpajnic](https://github.com/rokpajnic) - Idea & Initial work

See also the list of [contributors](https://github.com/GateHubNet/data-api-go/contributors)
who participated in this project.
