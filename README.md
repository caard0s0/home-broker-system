<div id="top"></div>

<!-- About the Project -->
<div align="center">
<h2>Home Broker Algorithm - Microservice</h2>
<p>A Home Broker algorithm for Investment Brokers, developed with <a href="https://go.dev/">Go</a>.</p>
<a href="https://github.com/caard0s0/home-broker-system/issues">Report Bugs</a>
&nbsp;&bull;&nbsp;
<a href="https://github.com/caard0s0/home-broker-system/actions">Actions</a>
&nbsp;&bull;&nbsp;
<a href="https://github.com/caard0s0/home-broker-system/pulls">Pull Requests</a>
</div>

&nbsp;

![home_broker_flowchart](https://github.com/caard0s0/home-broker-system/assets/95318788/107c4c7b-e64a-45c0-af85-4fc8a494edf2)

A Financial Software specialized in the intermediation of money between savers and those in need of loans, as well as in the custody of that money. It was created following SOLID principles, for better scalability and code maintenance. In addition, thinking about a reliable and well-tested application, with Unit and Automated Tests using Mock DB, the tests apply the concept of DB Stubs. Deploying it using Amazon's Cloud services.

&nbsp;

<h3>Built With</h3>

[![Tech Tools](https://skillicons.dev/icons?i=go,docker,kafka)](https://skillicons.dev)


<!-- Table of Contents -->
<details>
<summary>Table of Contents</summary>
<ol>
<li>
    <a href="#getting-started">Getting Started</a>
    <ul>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#usage">Usage</a></li>
        <li><a href="#tests">Tests</a></li>
    </ul>
</li>
<li><a href="#license">License</a></li>
<li><a href="#contact">Contact</a></li>
</ol>
</details>

&nbsp;


<!-- Getting Started -->
<h2 id="getting-started">Getting Started</h2>

<p>To get started, You need to have <strong>Go 1.21+</strong> installed on your machine, for more information visit <a href="https://go.dev/dl/">Go Downloads</a>. You also need to have <strong>Docker Desktop</strong> installed, for more information visit <a href="https://docs.docker.com/engine/install/">Docker Engine Install</a>.</p>

<p><strong>OBS:</strong> This guide is designed to run this project locally (Local Development), on Linux-based systems.</p>


<!-- Installation -->
<h3 id="installation">Installation</h3>

1. Clone the repository.
```bash
git clone https://github.com/caard0s0/home-broker-system.git
```


<!-- Usage -->
<h2 id="usage">Usage</h2>

<p>After completing the installation, you can run the project.</p>

1. Create and run the <strong>Containers</strong>.

```cmd
docker compose up -d
```

2. Open a browser tab at `localhost:9021` to access your <strong>Cluster</strong> in the <strong>Confluent Control Center</strong>.

![cluster_control_center](https://github.com/caard0s0/home-broker-system/assets/95318788/9360c92d-06cb-4b80-97f7-9e3f0cbdcc45)

3. Create a new <strong>Topic</strong>.

    <strong>WARNING:</strong> The <strong>Topic name</strong> must be exactly as written in the image.

![new_topic](https://github.com/caard0s0/home-broker-system/assets/95318788/42f84f86-b7b0-4385-8d76-e4f0dcc22427)

<!-- Tests -->
<h2 id="tests">Tests</h2>

<p>To be able to run all the tests, follow the command below.</p>

1. Run all the <strong>Tests</strong>.

```cmd
go test -v -cover ./...
```


<br>


<!-- License -->
<h2 id="license">License</h2>

This project is being distributed under the <strong>MIT License</strong>, see ```LICENSE.txt``` for more information.


<br>


<!-- Contact -->
<h2 id="contact">Contact</h2>

* Software Engineer  
* Vinicius Cardoso - <a href="mailto:cardoso.business.ctt@gmail.com">Email</a>

<p align="right">
<a href="#top"> &uarr; back to top</a>
</p> 
