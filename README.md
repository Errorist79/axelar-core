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

## Yerel olarak ikili dosyalar oluşturma

Doğrulayıcı düğüm için yerel ikili dosyalar oluşturmak üzere `make build` komutunu yürütün. Bunlar 'bin' klasöründe oluşturulur.

## Docker imajları oluşturma

Düğüm için normal bir docker görüntüsü oluşturmak için `make docker-image` komutunu çalıştırın. Bu, axelar/core:latest görüntüsünü oluşturur.

Hata ayıklama için bir docker görüntüsü oluşturmak için ([delve ile](https://github.com/go-delve/delve)) `make docker-image-debug` komutunu yürütün. Bu, axelar/core-debug:latest görüntüsünü oluşturur.

## Yerel bir düğümle etkileşim kurma

Yerel (dockerize edilmiş) bir düğüm çalışırken, düğümle etkileşim kurmak için `axelard` ikili dosyası kullanılabilir. Mevcut komutlar hakkında bilgi almak için ikili dosyaları oluşturduktan sonra `./bin/axelard --help` komutunu çalıştırın.

## API belgelerini göster

`Godoc`'un ana bilgisayara kurulu olduğundan emin olmak için `GO111MODULE=off go get -u golang.org/x/tools/cmd/godoc` dosyasını yürütün.

Kurulumdan sonra, yerel bir godoc sunucusunu barındırmak için `godoc -http ":{port}" -index` komutunu çalıştırın. Örneğin, `8080` numaralı bağlantı noktasıyla belgeler http://localhost:8080/pkg/github.com/axelarnetwork/axelar-core adresinde barındırılır. Dizin bayrağı, belgeleri aranabilir hale getirir.

Paketlerin başındaki, türlerden önceki ve işlevlerden önceki yorumlar, belgeleri doldurmak için kaynak dosyalardan otomatik olarak alınır. Daha fazla bilgi için https://blog.golang.org/godoc adresine bakın.

### CLI komut belgeleri

`axelard` için mevcut CLI komutlarının tam listesi için [buraya](docs/cli/toc.md) bakın 

## Test araçları

Yürütülebilir bir dosya olduğundan, ``go mod download`` veya benzeri komutlar yürütülürken otomatik olarak indirilmez. Arayüzler için mocklar oluşturmak üzere moq aracını kurmak için ``go get github.com/matryer/moq`` komutunu çalıştırın.

## Hata ödülü ve güvenlik açıklarının açıklanması

Bakın; [Axelar dökümanları](https://docs.axelar.dev/#/bug-bounty).
