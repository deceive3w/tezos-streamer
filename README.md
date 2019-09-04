# tezos-streamer is an experimental middleware for streaming Tezos data

[![CircleCI](https://circleci.com/gh/ecadlabs/tezos-streamer.svg?style=svg)](https://circleci.com/gh/ecadlabs/tezos-streamer)

_WARNING: This project is in early stage, active development. While we welcome users and feedback, please be warned that this project is a work in progress and users should proceed with caution._

## What is tezos-streamer?

`tezos-streamer` offers a websocket interface that allows callers to "subscribe" to certain operations on the Tezos Blockchain.

It is essentially middleware that sits between a Tezos Node's RPC interface and a websocket client (such as a web browser). Clients can "subscribe" to an address, and expect to have balance or storage updates streamed via the websocket efficiently and promptly.

It is our hope that this experimental websocket API will help inform a future, similar API directly in the Tezos RPC node.

## Getting started

Docker images and pre-built binaries are available from the [releases](https://github.com/ecadlabs/tezos-streamer/releases) github page.

## Reporting Issues

### Security Issues

To report a security issue, please contact security@ecadlabs.com or via [keybase/jevonearth][1] on keybase.io.

Reports may be encrypted using keys published on keybase.io using [keybase/jevonearth][1].

### Other Issues & Feature Requests

Please use the [GitHub issue tracker](https://github.com/ecadlabs/tezos-streamer/issues) to report bugs or request features.

## Contributions

To contribute, please check the issue tracker to see if an existing issue exists for your planned contribution. If there's no Issue, please create one first, and then submit a pull request with your contribution.

For a contribution to be merged, it must be well documented, come with unit tests, and integration tests where appropriate. Submitting a "work in progress" pull request is welcome!

---

## Disclaimer

THIS SOFTWARE IS PROVIDED "AS IS" AND ANY EXPRESSED OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

[1]: https://keybase.io/jevonearth
