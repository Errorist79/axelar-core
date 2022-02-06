⚠️⚠️⚠️ **BU DEVAM EDEN BİR ÇALIŞMADIR** ⚠️⚠️⚠️

# axelar-core

Cosmos SDK'yı temel alan axelar-core uygulaması, axelar ağının ana uygulamasıdır. Bu depo, bir çekirdek düğümü çalıştırmak için gerekli ikili dosyaları ve Docker imajı oluşturmak için kullanılır.

## İkili dosyalar ve Docker görüntüleri oluşturmak için ön koşullar

1. Makinenizde bir SSH anahtarınız olsun
2. Kimlik doğrulama için genel anahtarınızı Github hesabınıza ekleyin
3. Özel anahtarınızı ssh aracınıza eklemek için `ssh-add ~/.ssh/{private key file name}` çalıştırın. **ÖNEMLİ**: ssh aracısı yalnızca özel anahtarınızı bellekte tutar, bu nedenle makinenizi her yeniden başlattığınızda bu adımı tekrarlamanız gerekir. Bu adımı [burada](https://apple.stackexchange.com/questions/254468/macos-sierra-doesn-t-seem-to-remember-ssh-keys-between-reboots/264974#264974) açıklandığı gibi iki şekilde otomatikleştirebilirsiniz :
    * `~/.ssh/config` Dosyanıza şunları ekleyin:
    ```
    Host *
       AddKeysToAgent yes
       UseKeychain yes     
    ```
    * `ssh-add ~/.ssh/{private key file name} &>/dev/null` parametresini Kabuğunuzun .rc dosyasına ekleyin (örn. `~/.bash_profile`).
4. `go get` zorlamak için ssh ile kimlik doğrulaması yapın; `git config --global url."git@github.com:axelarnetwork".insteadOf https://github.com/axelarnetwork`

## Building binaries locally

Execute `make build` to create local binaries for the validator node. They are created in the `./bin` folder.

## Creating docker images

To create a regular docker image for the node, execute `make docker-image`. This creates the image axelar/core:
latest.

To create a docker image for debugging (with [delve](https://github.com/go-delve/delve)),
execute `make docker-image-debug`. This creates the image axelar/core-debug:latest.

## Interacting with a local node

With a local (dockerized) node running, the `axelard` binary can be used to interact with the node.
Run `./bin/axelard --help` after building the binaries to get information about the available commands.

## Show API documentation

Execute `GO111MODULE=off go get -u golang.org/x/tools/cmd/godoc` to ensure that `godoc` is installed on the host.

After the installation, execute `godoc -http ":{port}" -index` to host a local godoc server. For example, with
port `8080` the documentation is hosted at
http://localhost:8080/pkg/github.com/axelarnetwork/axelar-core. The index flag makes the documentation searchable.

Comments at the beginning of packages, before types and before functions are automatically taken from the source files
to populate the documentation. See https://blog.golang.org/godoc for more information.

### CLI command documentation

For the full list of available CLI commands for `axelard` see [here](docs/cli/toc.md)

## Test tools

Because it is an executable, github.com/matryer/moq is not automatically downloaded when executing ``go mod download``
or similar commands. Execute ``go get github.com/matryer/moq`` to install the _moq_ tool to generate mocks for
interfaces.

## Bug bounty and disclosure of vulnerabilities

See the [Axelar documentation website](https://docs.axelar.dev/#/bug-bounty).
