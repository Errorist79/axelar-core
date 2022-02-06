# CLI komutuna genel bakış 

- [axelard](axelard.md) - Axelar Uygulaması 
  - [add-genesis-account \[address_or_key_name\] \[coin\]\[,\[coin\]\]](axelard_add-genesis-account .md) - Genesis.json'a bir genesis hesabı ekle 
  - [add-genesis-evm-chain \[name\] \[native entity\]](axelard_add-genesis-evm-chain.md) - Bir EVM zinciri ekler genesis.json 
  - [collect-gentxs](axelard_collect-gentxs.md) - Genesis tx'lerini toplayın ve bir genesis.json dosyası çıktısı alın 
  - [debug](axelard_debug.md) - Uygulamanızda hata ayıklamaya yardımcı olan araç 
    - [addr \[address \]](axelard_debug_addr.md) - Bir adresi hex ve bech32 arasında dönüştürün 
    - [pubkey \[pubkey\]](axelard_debug_pubkey.md) - Proto JSON'dan bir pubkey kodunu çözün
    - [raw-bytes \[raw-bytes\]](axelard_debug_raw-bytes.md) - Ham bayt çıktısını (örn. \[10 21 13 255\]) 
  hex'e dönüştürün - [export](axelard_export.md) - Dışa Aktar durumu 
  JSON'a - [gentx \[anahtar_adı\] \[amount\]](axelard_gentx.md) - Kendi kendine yetki veren bir genesis tx oluşturun 
  - [sağlık kontrolü](axelard_health-check.md) - 
  - [init \[ moniker\]](axelard_init.md) - Özel doğrulayıcı, p2p, genesis ve uygulama yapılandırma dosyalarını 
  başlat - [keys](axelard_keys.md) - Uygulamanızın anahtarlarını yönetin 
    - [add \<name>](axelard_keys_add.md) - Şifrelenmiş bir özel anahtar (yeni oluşturulmuş veya kurtarılmış) ekleyin, şifreleyin ve <name> dosyasına kaydedin 
    - [delete \<name>...](axelard_keys_delete.md) - Verilen anahtarları silin
    - [dışa aktarma \<ad>](axelard_keys_export.md) - Özel anahtarları dışa 
    aktarma - [içe aktarma \<ad> \<keyfile>](axelard_keys_import.md) - Özel anahtarları yerel anahtar tabanına içe aktarma 
    - [list](axelard_keys_list.md ) - Tüm anahtarları listele 
    - [migrate \<old_home_dir>](axelard_keys_migrate.md) - Anahtarları eski (db tabanlı) Keybase'den taşıyın 
    - [mnemonic](axelard_keys_mnemonic.md) - Bazı giriş entropileri için bip39 anımsatıcısını hesaplayın 
    - [ parse \<hex-or-bech32-address>](axelard_keys_parse.md) - Adresi hex'ten bech32'ye ve tam tersi şekilde ayrıştırın 
    - [show \[name_or_address \[name_or_address...\]\]](axelard_keys_show.md) - Anahtar bilgileri ada veya adrese göre alın
  - [migrate \[hedef-version\] \[genesis-file\]](axelard_migrate.md) - Oluşumu belirli bir hedef sürüme taşıyın 
  - [sorgu](axelard_query.md) - Alt komutları sorgulama 
    - [hesap \[adres\ ]](axelard_query_account.md) - Adrese göre hesap 
    sorgulama - [auth](axelard_query_auth.md) - Yetkilendirme modülü için sorgulama komutları 
      - [account \[address\]](axelard_query_auth_account.md) - Adrese göre hesap sorgulama 
      - [accounts](axelard_query_auth_accounts.md) - Tüm hesapları sorgulayın 
      - [params](axelard_query_auth_params.md) - Mevcut auth parametrelerini 
    sorgulayın - [bank](axelard_query_bank.md) - Banka modülü için sorgulama komutları
      - [balances \[address\]](axelard_query_bank_balances.md) - Adrese göre hesap bakiyesi sorgulama 
      - [denom-metadata](axelard_query_bank_denom-metadata.md) - Madeni para birimleri için müşteri meta verilerini 
      sorgulama - [total](axelard_query_bank_ ) - Zincirin toplam madeni para arzını sorgulayın 
    - [blok \[yükseklik\]](axelard_query_block.md) - Belirli bir yükseklikte bir blok için doğrulanmış verileri alın 
    - [distribution](axelard_query_distribution.md) - dağıtım modülü 
      - [commission \[validator\]](axelard_query_distribution_commission.md) - Dağıtım doğrulayıcı komisyonunu sorgulayın 
      - [topluluk havuzu](axelard_query_distribution_community-pool.md) - Topluluk havuzundaki jeton miktarını sorgulayın
      - [params](axelard_query_distribution_params.md) - Dağıtım parametrelerini sorgulayın - [rewards 
      \[delegator-addr\] \[validator-addr\]](axelard_query_distribution_rewards.md) - Belirli bir doğrulayıcıdan tüm dağıtım delege edici ödüllerini veya ödüllerini sorgulayın 
      - [ eğik çizgi \[doğrulayıcı\] \[başlangıç ​​yüksekliği\] \[bitiş yüksekliği\]](axelard_query_distribution_slashes.md) - Sorgu dağıtım doğrulayıcı eğik çizgileri 
      - [doğrulayıcı-olağanüstü-rewards \[doğrulayıcı\]](axelard_query_distribution_validator-olağanüstü-ödüller .md) - Bir doğrulayıcı ve tüm delegasyonları için bekleyen (geri 
    alınmamış) ödülleri sorgulayın - [kanıt](axelard_query_evidence.md) - Karma veya tüm (sayfalandırılmış) gönderilen kanıtlar için kanıt sorgulayın
    - [evm](axelard_query_evm.md) - evm modülü için komutları sorgulama 
      - [address \[chain\]](axelard_query_evm_address.md) - EVM adresini döndürür 
      - [batch-commands \[chain\] \[batchCommandsID\] ](axelard_query_evm_batched-commands.md) - Axelar Gateway'de yürütülecek bir EVM işlemine sarılabilecek imzalı toplu komutları alın 
      - [yazıcı bilgisi \[depozito adresi\]](axelard_query_evm_burner-info.md) - Bilgi alın bir yazıcı adresi hakkında 
      - [bytecode \[chain\] \[contract\]](axelard_query_evm_bytecode.md) - Zincir \[chain\] için bir EVM sözleşmesinin \[contract\] bayt 
      kodlarını getirin - [chains](axelard_query_evm_chains. md) - EVM zincirlerini al
      - [komut \[chain\] \[id\]](axelard_query_evm_command.md) - Bir zincir ve komut kimliği verilen bir EVM ağ geçidi komutu hakkında bilgi alın 
      - [confirmation-height \[chain\]](axelard_query_evm_confirmation-height. md) - Verilen zincir için minimum onay yüksekliğini döndürür 
      - [depozit-state \[chain\] \[txID\] \[burner address\] \[amount\]](axelard_query_evm_deposit-state.md) - Durumu sorgulayın - [gateway-address \[chain\]]( axelard_query_evm_gateway-address.md 
      ) - Axelar Gateway sözleşme adresini sorgulayın 
      - [latest-batch-commands \[chain\]](axelard_query_evm_latest-bached-commands.md ) - Axelar Gateway'de yürütülecek bir EVM işlemine sarılabilecek en son toplu komutları alın
      - [pending-commands \[chain\]](axelard_query_evm_pending-commands.md) - Henüz bir gruba eklenmemiş komutların listesini alın 
      - [token-address \[chain\]](axelard_query_evm_token-address.md) - Sorgu sembol veya varlığa göre bir belirteç adresi 
    - [feegrant](axelard_query_feegrant.md) - Fegrant modülü için sorgulama komutları 
      - [grant \[granter\] \[grantee\]](axelard_query_feegrant_grant.md) - Tek bir öğenin ayrıntılarını sorgulayın hibe 
      - [grants \[grantee\]](axelard_query_feegrant_grants.md) - Bir hibe sahibinin tüm hibelerini 
    sorgulama - [gov](axelard_query_gov.md) - Yönetim modülü için sorgulama komutları 
      - [deposit \[teklif-kimliği\] \[ depoziter-addr\]](axelard_query_gov_deposit.md) - Depozito ayrıntılarını sorgula
      - [deposits \[proposal-id\]](axelard_query_gov_deposits.md) - Bir teklif üzerindeki mevduatları sorgulayın 
      - [param \[param-type\]](axelard_query_gov_param.md) - Parametrelerini sorgulayın (oylama|tallying|deposit) yönetim süreci 
      - [params](axelard_query_gov_params.md) - Yönetim sürecinin parametrelerini sorgulayın 
      - [teklif \[teklif-kimliği\]](axelard_query_gov_proposal.md) - Tek bir teklifin ayrıntılarını sorgulayın 
      - [teklifler](axelard_query_gov_proposals. md) - İsteğe bağlı filtrelerle teklifleri sorgulayın 
      - [teklif eden \[teklif-kimliği\]](axelard_query_gov_proposer.md) - Bir yönetim teklifi teklif edeni sorgulayın 
      - [tally \[teklif-kimliği\]](axelard_query_gov_tally.md) - Al bir teklif oylamasının çetelesi
      - [oy \[teklif-kimliği\] \[voter-addr\]](axelard_query_gov_vote.md) - Tek bir oylamanın ayrıntılarını sorgulayın 
      - [oylar \[teklif-kimliği\]](axelard_query_gov_votes.md) - Üzerindeki oyları sorgulayın bir teklif 
    - [ibc](axelard_query_ibc.md) - IBC modülü için sorgulama komutları 
      - [channel](axelard_query_ibc_channel.md) - IBC kanal sorgusu alt komutları 
        - [channels](axelard_query_ibc_channel_channels.md) - Tüm kanalları sorgula 
        - [client-state \[port-id\] \[channel-id\]](axelard_query_ibc_channel_client-state.md) - Bir kanalla ilişkili istemci durumunu sorgulayın 
        - [bağlantılar \[bağlantı kimliği\]](axelard_query_ibc_channel_connections.md) - Tümünü sorgulayın bir bağlantıyla ilişkili kanallar
        - [end \[port-id\] \[channel-id\]](axelard_query_ibc_channel_end.md) - Bir kanal sonunu sorgulayın 
        - [sonraki dizi-alıcı \[port-id\] \[channel-id\]] (axelard_query_ibc_channel_next-sequence-receive.md) - Bir sonraki alma sırasını sorgulayın 
        - [packet-ack \[port-id\] \[channel-id\] \[sequence\]](axelard_query_ibc_channel_packet-ack.md) - Sorgu a paket onayı 
        - [paket taahhüt \[bağlantı noktası kimliği\] \[kanal kimliği\] \[sıralama\]](axelard_query_ibc_channel_packet-commitment.md) - Bir paket taahhüdü sorgulama 
        - [paket taahhütleri \[bağlantı noktası kimliği\ ] \[channel-id\]](axelard_query_ibc_channel_packet-commitments.md) - Bir kanalla ilişkili tüm paket taahhütlerini sorgulayın
        - [paket makbuzu \[bağlantı noktası kimliği\] \[kanal kimliği\] \[sıra\]](axelard_query_ibc_channel_packet-receipt.md) - Bir paket makbuzu 
        sorgulayın - [alınmayan onaylar \[bağlantı noktası kimliği\] \ [channel-id\]](axelard_query_ibc_channel_unreceived-acks.md) - Bir kanalla ilişkili tüm alınmamış onayları sorgulayın 
        - [unrecepted-packets \[port-id\] \[channel-id\]](axelard_query_ibc_channel_unreceived-packets.md ) - Bir kanalla ilişkili tüm alınmamış paketleri 
      sorgulayın - [client](axelard_query_ibc_client.md) - IBC istemci sorgu alt komutları 
        - [consensus-state \[client-id\] \[height\]](axelard_query_ibc_client_consensus-state.md) - Belirli bir yükseklikte bir müşterinin fikir birliği durumunu sorgulayın
        - [consensus-states \[client-id\]](axelard_query_ibc_client_consensus-states.md) - Bir istemcinin tüm fikir birliği durumlarını sorgulayın. 
        - [header](axelard_query_ibc_client_header.md) - Çalışan zincirin en son başlığını 
        sorgulayın - [params](axelard_query_ibc_client_params.md) - Mevcut ibc istemci parametrelerini sorgulayın 
        - [self-consensus-state](axelard_query_ibc_mss-state.self) ) - Bu zincir için kendi konsensüs durumunu sorgulayın 
        - [durum \[client-id\]](axelard_query_ibc_client_state.md) - İstemci durumunu 
        sorgulayın - [durumlar](axelard_query_ibc_client_states.md) - Mevcut tüm hafif istemcileri sorgulayın 
        - [status \ [client-id\]](axelard_query_ibc_client_status.md) - İstemci durumunu sorgula
      - [bağlantı](axelard_query_ibc_connection.md) - IBC bağlantı sorgusu alt komutları 
        - [bağlantılar](axelard_query_ibc_connection_connections.md) - Tüm bağlantıları sorgula 
        - [end \[bağlantı kimliği\]](axelard_query_ibc_connection_end.md) 
        - [ Depolanan bağlantıyı sorgula yol \[client-id\]](axelard_query_ibc_connection_path.md) - Depolanan istemci bağlantı yollarını sorgulayın 
    - [ibc-transfer](axelard_query_ibc-transfer.md) - IBC uygun token aktarım sorgusu alt komutları 
      - [denom-trace \[hash\] ](axelard_query_ibc-transfer_denom-trace.md) - Belirli bir izleme karma 
      değerinden denom izleme bilgisini sorgula - [denom-traces](axelard_query_ibc-transfer_denom-traces.md) - Tüm simge değerleri için izleme bilgilerini sorgula
      - [escrow-address](axelard_query_ibc-transfer_escrow-address.md) - Bir kanalın emanet adresini alın 
      - [params](axelard_query_ibc-transfer_params.md) - Geçerli ibc-transfer parametrelerini 
    sorgulayın - [mint](axelard_md_mint) ) - Darphane modülü için sorgulama komutları 
      - [yıllık-provisions](axelard_query_mint_annual-provisions.md) - Mevcut basım yıllık provizyon değerini 
      sorgulayın - [inflation](axelard_query_mint_inflation.md) - Mevcut basım enflasyon değerini sorgulayın 
      - [params]( axelard_query_mint_params.md) - Geçerli basım parametrelerini sorgula 
    - [nexus](axelard_query_nexus.md) - nexus modülü için komutları sorgulama
      - [chain-maintainers \[chain\]](axelard_query_nexus_chain-maintainers.md) - Verilen zincir için zincir bakımcılarını döndürür 
      - [en son-deposit-address \[deposit chain\] \[alıcı zinciri\] \[alıcı adresi \]](axelard_query_nexus_latest-deposit-address.md) - Adrese göre hesap sorgusu 
      - [zincir için transferler \[chain\] \[durum (beklemede|arşivlendi)\]](axelard_query_nexus_transfers-for-chain.md) - Adrese göre hesap 
    sorgulama - [params](axelard_query_params.md) - Params modülü için sorgulama komutları 
      - [subspace \[subspace\] \[key\]](axelard_query_params_subspace.md) - Altuzay ve anahtara göre ham parametreleri sorgulama 
    - [izin](axelard_query_permission.md) - İzin modülü için komutları sorgulama
      - [yönetişim anahtarı](axelard_query_permission_governance-key.md) - Yönetişim anahtarını döndürür 
    - [eğik çizgi](axelard_query_slashing.md) - Eğik modül için sorgulama komutları 
      - [params](axelard_query_slashing_params.md - geçerli kesme parametrelerini sorgulama) 
      - [signing-info \[validator-conspub\]](axelard_query_slashing_signing-info.md) - Doğrulayıcının imzalama bilgilerini 
      sorgulayın - [signing-infos](axelard_query_slashing_signing-infos.md) - Tüm doğrulayıcıların imzalama bilgilerini sorgulayın 
    - [anlık görüntü]( axelard_query_snapshot.md) - Anlık görüntü modülü için komutları sorgulama 
      - [info \[counter\]](axelard_query_snapshot_info.md) - Belirli bir sayaç için anlık görüntüyü getirir
      - [operatör \[proxy adresi\]](axelard_query_snapshot_operator.md) - \[proxy address\] ile ilişkili operatör adresini alın 
      - [proxy \[operatör adresi\]](axelard_query_snapshot_proxy.md) - İle ilişkili proxy adresini getirin \[operatör adresi\] ve durum (etkin/etkin değil) 
      - [doğrulayıcılar](axelard_query_snapshot_validators.md) - Doğrulayıcıları ve bilgilerini 
    alın - [stake](axelard_query_staking.md) - Stake modülü için sorgulama komutları 
      - [delegasyon \[ delegator-addr\] \[validator-addr\]](axelard_query_staking_delegation.md) - Adrese ve doğrulayıcı adresine dayalı olarak yetki sorgulama 
      - [delegations \[delegator-addr\]](axelard_query_staking_delegations.md) - Tarafından yapılan tüm yetkilendirmeleri sorgula bir delege
      - [delegations-to \[validator-addr\]](axelard_query_staking_delegations-to.md) - Bir doğrulayıcıya yapılan tüm yetkilendirmeleri sorgulayın 
      - [historical-info \[height\]](axelard_query_staking_historical-info.md) - Geçmiş bilgileri sorgulayın belirli bir yükseklikte 
      - [params](axelard_query_staking_params.md) - Mevcut stake etme parametreleri bilgilerini 
      sorgulayın - [pool](axelard_query_staking_pool.md) - Mevcut stake havuzu değerlerini sorgulayın 
      - [redelegation \[delegator-addr\] \[src-validator -addr\] \[dst-validator-addr\]](axelard_query_staking_redelegation.md) - Yetki verene ve kaynak ve hedef doğrulayıcı adresine dayalı bir yeniden yetkilendirme kaydı sorgulama
      - [redelegations \[delegator-addr\]](axelard_query_staking_redelegations.md) - Bir yetkilendiren için tüm yeniden yetkilendirme kayıtlarını sorgula 
      - [redelegations-from \[validator-addr\]](axelard_query_staking_redelegations-from.md) - Giden tüm redelegatation'ları sorgula a validator 
      - [unbonding-delegation \[delegator-addr\] \[validator-addr\]](axelard_query_staking_unbonding-delegation.md) - Delegator ve validator adresine dayalı bir unbonding-delegation kaydını sorgulayın 
      - [unbonding-delegations \[delegator -addr\]](axelard_query_staking_unbonding-delegations.md) - Bir temsilci için tüm bağlayıcı olmayan yetkilendirme kayıtlarını sorgula
      - [unbonding-delegations-from \[validator-addr\]](axelard_query_staking_unbonding-delegations-from.md) - Bir doğrulayıcıdan tüm bağlayıcı olmayan yetkilendirmeleri sorgulayın 
      - [validator \[validator-addr\]](axelard_query_staking_validator.md) - a validator 
      - [validators](axelard_query_staking_validators.md) - Tüm doğrulayıcılar için sorgu 
    - [tendermint-validator-set \[height\]](axelard_query_tendermint-validator-set.md) - Verilen yükseklikte tam ihalemint doğrulayıcı setini alın 
    - [ tss](axelard_query_tss.md) - tss modülü için komutları sorgulama 
      - [aktif-eski-anahtarlar \[zincir\] \[role\]](axelard_query_tss_active-old-keys.md) - Doğrulayıcı tarafından aktif eski anahtar kimliklerini sorgulayın
      - [aktif-eski-anahtarlar-doğrulayıcı \[doğrulayıcı adresi\]](axelard_query_tss_active-old-keys-by-validator.md) - Doğrulayıcı tarafından aktif eski anahtar kimliklerini sorgulayın 
      - [deaktive edilmiş operatörler](axelard_query_tss_deactivated-operators. md) - Devre dışı bırakılan operatör adreslerinin listesini 
      getir - [external-key-id \[chain\]](axelard_query_tss_external-key-id.md) - Verilen zincir için mevcut harici anahtarların anahtar kimliklerini döndürür 
      - [anahtar \ [anahtar kimliği\]](axelard_query_tss_key.md) - Anahtar kimliğine göre bir anahtar sorgulayın 
      - [key-id \[chain\] \[role\]](axelard_query_tss_key-id.md) - keyChain ve keyRole kullanarak anahtar kimliğini sorgulayın 
      - [key-shares-by-key-id \[anahtar kimliği\]](axelard_query_tss_key-shares-by-key-id.md) - Sorgu anahtarı, bilgileri anahtar kimliğine göre paylaşır
      - [doğrulayıcıya göre anahtar paylaşımları \[doğrulayıcı adresi\]](axelard_query_tss_key-shares-by-validator.md) - Doğrulayıcıya göre sorgu anahtarı paylaşım bilgileri 
      - [sonraki anahtar kimliği \[zincir\] \[rol\ ]](axelard_query_tss_next-key-id.md) - Belirli bir zincirdeki sonraki dönüş ve verilen anahtar rolü için atanan anahtar kimliğini döndürür 
      - [kurtarma \[doğrulayıcı adresi\] \[anahtar kimliği #1\] .. . \[anahtar kimliği #N\]](axelard_query_tss_recover.md) - Belirtilen anahtar kimliği için paylaşımları kurtarma girişimi 
      - [signature \[sig ID\]](axelard_query_tss_signature.md) - Sig kimliğine göre bir imza sorgulama 
    - [ tx --type=\[hash|acc_seq|signature\] \[hash|acc_seq|signature\]](axelard_query_tx.md) - Bir işlem için hash, "<addr>/<seq>" kombinasyonu veya virgül ile sorgu- taahhüt edilmiş bir blokta ayrılmış imzalar
    - [txs](axelard_query_txs.md) - Bir dizi olayla eşleşen sayfalandırılmış işlemler için sorgu 
    - [yükseltme](axelard_query_upgrade.md) - Yükseltme modülü için sorgulama komutları 
      - [uygulandı \[yükseltme-adı\]](axelard_query_upgrade_applied. md) - tamamlanmış bir yükseltmenin uygulandığı yükseklik için blok başlığı 
      - [module_versions \[opsiyonel modül_adı\]](axelard_query_upgrade_module_versions.md) - modül sürümlerinin listesini alın 
      - [plan](axelard_query_upgrade_plan.md) - yükseltme planını alın ( varsa) 
  - [rosetta](axelard_rosetta.md) - bir rosetta sunucusunu çalıştırın 
  - [set-genesis-auth](axelard_set-genesis-auth.md) - Yetkilendirme modülü için genesis parametrelerini ayarlayın
  - [set-genesis-chain-params \[bitcoin | evm\] \[chain\]](axelard_set-genesis-chain-params.md) - Genesis.json'da zincir parametrelerini ayarlayın 
  - [set-genesis-crisis](axelard_set-genesis-crisis.md) - Genesis parametrelerini ayarlayın kriz modülü için 
  - [set-genesis-evm-contracts](axelard_set-genesis-evm-contracts.md) - EVM'nin sözleşme parametrelerini genesis.json 
  - [set-genesis-gov](axelard_set-genesis-gov. md) - Yönetim modülü için genesis parametrelerini ayarlayın 
  - [set-genesis-mint](axelard_set-genesis-mint.md) - Mint modülü için genesis parametrelerini ayarlayın 
  - [set-genesis-reward](axelard_set-genesis- ödül.md) - Ödül modülü için oluşum parametrelerini ayarlayın
  - [set-genesis-slashing](axelard_set-genesis-slashing.md) - Slashing modülü için genesis parametrelerini ayarlayın 
  - [set-genesis-snapshot](axelard_set-genesis-snapshot.md) - Genesis parametrelerini ayarlayın anlık görüntü modülü 
  - [set-genesis-staking](axelard_set-genesis-staking.md) - Stake modülü için genesis parametrelerini ayarlayın 
  - [set-genesis-tss](axelard_set-genesis-tss.md) - Genesis parametrelerini ayarlayın tss modülü için 
  - [set-genesis-vote](axelard_set-genesis-vote.md) - Oylama modülü için genesis parametrelerini ayarlayın 
  - [set-governance-key \[threshold\] \[\[pubKey\]. ..\]](axelard_set-governance-key.md) - axelar ağı için genesis multisig yönetişim anahtarını ayarla 
  - [start](axelard_start.md) - Tam düğümü çalıştır
  - [status](axelard_status.md) - Uzak düğümü durum için sorgula 
  - [tendermint](axelard_tendermint.md) - Tendermint alt komutları 
    - [show-address](axelard_tendermint_show-address.md) - Bu düğümün ihalemint doğrulayıcı mutabakat adresini gösterir 
    - [ show-node-id](axelard_tendermint_show-node-id.md) - Bu düğümün kimliğini göster 
    - [show-validator](axelard_tendermint_show-validator.md) - Bu düğümün ihalemint doğrulayıcı bilgilerini göster 
    - [sürüm](axelard_tendermint_version.md) - İhale kütüphanelerinin sürümünü yazdırın 
  - [tx](axelard_tx.md) - İşlemler alt komutları 
    - [axelarnet](axelard_tx_axelarnet.md) - axelarnet işlemleri alt komutları
      - [add-cosmos-based-chain \[name\] \[yerel varlık\] \[min miktarı\] \[adres öneki\]](axelard_tx_axelarnet_add-cosmos-based-chain.md) - Yeni bir kozmos tabanlı ekleyin zincir 
      - [depozitoyu onayla \[denom\] \[burnerAddr\]](axelard_tx_axelarnet_confirm-deposit.md) - Varlık değeri ve brülör adresi ile gönderilen Axelar zincirine bir depozito onaylayın 
      - [yürütme-beklemede-transferler]( axelard_tx_axelarnet_execute-pending-transfers.md) - Bekleyen tüm aktarımları Axelar zincirine gönderin 
      - [link \[alıcı zinciri\] \[alıcı adresi\] \[varlık\]](axelard_tx_axelarnet_link.md) - Bir çapraz zincir adresini bir Axelar address 
      - [register-asset \[chain\] \[denom\] \[min miktar\]](axelard_tx_axelarnet_register-asset.md) - Kozmos tabanlı bir zincire yeni bir varlık kaydedin
      - [kayıt-ücret-toplayıcı \[ücret toplayıcı\]](axelard_tx_axelarnet_register-fee-collector.md) - axelarnet ücret tahsildarı hesabını kaydedin 
      - [kayıt-yolu \[zincir\] \[yol\]](axelard_tx_axelarnet_kayıt yolu. md) - Bir kozmos zinciri için bir ibc yolu 
      kaydedin - [route-ibc-transfers](axelard_tx_axelarnet_route-ibc-transfers.md) - Kozmos zincirlerine transferleri bekleyen rotalar 
    - [banka](axelard_tx_bank.md) - Banka işlemi alt komutları 
      - [ send \[from_key_or_address\] \[to_address\] \[amount\]](axelard_tx_bank_send.md) - Bir hesaptan diğerine para gönderin. 
        \[from_key_or_address\] öğesinden ima edildiği için '--from ' bayrağının yok sayıldığını unutmayın.
    - [yayın \[file_path\]](axelard_tx_broadcast.md) - Çevrimdışı oluşturulan yayın işlemleri 
    - [crisis](axelard_tx_crisis.md) - Kriz işlemleri alt komutları 
      - [değişmeyen bozuk \[modül adı\] \[değişmeyen rota\ ]](axelard_tx_crisis_invariant-broken.md) - Bir değişmezin zinciri durdurmak için kırıldığının kanıtını gönderin 
    - [decode \[amino-byte-string\]](axelard_tx_decode.md) - İkili kodlanmış bir işlem dizesinin kodunu çözün 
    - [dağıtım]( axelard_tx_distribution.md) - Dağıtım işlemleri alt komutları 
      - [fund-community-pool \[amount\]](axelard_tx_distribution_fund-community-pool.md) - Topluluk havuzuna belirtilen miktarda fon sağlar
      - [set-withdraw-addr \[withdraw-addr\]](axelard_tx_distribution_set-withdraw-addr.md) - bir adresle ilişkili ödüller için varsayılan para çekme adresini değiştirin 
      - [redraw-all-rewards](axelard_tx_distribution_withdraw-all-rewards .md) - bir yetki veren için tüm yetkilendirme ödüllerini geri 
      çekme - [redraw-rewards \[validator-addr\]](axelard_tx_distribution_withdraw-rewards.md) - Belirli bir yetkilendirme adresinden ödülleri geri çekme ve isteğe bağlı olarak yetkilendirme adresi verilmişse doğrulayıcı komisyonunu geri çekme bir doğrulayıcı operatördür 
    - [encode \[file\]](axelard_tx_encode.md) - Çevrimdışı oluşturulan işlemleri kodlayın 
    - [kanıt](axelard_tx_evidence.md) - Kanıt işlemi alt komutları 
    - [evm](axelard_tx_evm.md) - evm işlemleri alt komutları
      - [add-chain \[name\] \[yerel varlık\] \[anahtar türü\] \[chain config\]](axelard_tx_evm_add-chain.md) - Yeni bir EVM zinciri ekleyin 
      - [confirm-chain \[chain \]](axelard_tx_evm_confirm-chain.md) - Belirli bir ad ve yerel varlık için bir EVM zincirini onaylayın 
      - [confirm-erc20-deposit \[chain\] \[txID\] \[amount\] \[burnerAddr\]] (axelard_tx_evm_confirm-erc20-deposit.md) - Belirli miktarda jetonu bir yazıcı adresine gönderen bir EVM zincir işleminde bir ERC20 yatırmasını onaylayın 
      - [confirm-erc20-token \[chain\] \[origin chain\] \[origin varlık\] \[txID\]](axelard_tx_evm_confirm-erc20-token.md) - Bazı kaynak zinciri ve ağ geçidi adresinin belirli bir varlığı için bir EVM zinciri işleminde bir ERC20 belirteci dağıtımını onaylayın
      - [confirm-gateway-deployment \[chain\] \[txID\] \[address\]](axelard_tx_evm_confirm-gateway-deployment.md) - Ağ geçidi sözleşmesinin belirtilen zincir için belirtilen zincir için verilen adres 
      - [confirm-transfer-operatorship \[chain\] \[txID\] \[keyID\]](axelard_tx_evm_confirm-transfer-operatorship.md) - Bir EVM zincir işleminde bir transfer operatörlüğünü onaylayın 
      - [confirm-transfer- sahiplik \[chain\] \[txID\] \[keyID\]](axelard_tx_evm_confirm-transfer-ownership.md) - Bir EVM zincir işleminde bir transfer sahipliğini onaylayın 
      - [create-burn-tokens \[chain\]]( axelard_tx_evm_create-burn-tokens.md) - Bir EVM zincirinde onaylanmış tüm jeton mevduatları için yakma komutları oluşturun
      - [create-deploy-token \[evm chain\] \[başlangıç ​​zinciri\] \[kaynak varlık\] \[belirteç adı\] \[sembol\] \[ondalık sayılar\] \[kapasite\] \[min depozito \]](axelard_tx_evm_create-deploy-token.md) - AxelarGateway sözleşmesiyle bir konuşlandırma belirteci komutu 
      oluşturun - [create-pending-transfers \[chain\]](axelard_tx_evm_create-pending-transfers.md) - Tümünü işlemek için komutlar oluşturun bir EVM zincirine bekleyen aktarımlar 
      - [bağ \[zincir\] \[alıcı zinciri\] \[alıcı adresi\] \[varlık adı\]](axelard_tx_evm_link.md) - Bir çapraz zincir adresini oluşturulan bir EVM zinciri adresine bağlayın Axelar tarafından 
      - [sign-commands \[chain\]](axelard_tx_evm_sign-commands.md) - Bir EVM zincir sözleşmesi için bekleyen komutları imzalayın
      - [transfer-operatorship \[chain\] \[keyID\]](axelard_tx_evm_transfer-operatorship.md) - Bir EVM zincir sözleşmesi için transfer operatörlüğü komutu oluşturun 
      - [transfer-mülkiyet \[chain\] \[keyID\]]( axelard_tx_evm_transfer-ownership.md) - Bir EVM zincir sözleşmesi için devir sahipliği komutu oluşturun 
    - [feegrant](axelard_tx_feegrant.md) - Feegrant işlemleri alt komutları 
      - [grant \[granter_key_or_address\] \[grantee\]](axelard_tx_feegrant) - Bir adres için ücret ödeneği 
      - [revoke \[granter\] \[grantee\]](axelard_tx_feegrant_revoke.md) - ücret hibesini iptal et 
    - [gov](axelard_tx_gov.md) - Yönetim işlemleri alt komutları
      - [depozit \[teklif-kimliği\] \[deposit\]](axelard_tx_gov_deposit.md) - Etkin bir teklif için yatırılan jetonlar 
      - [teklif gönder](axelard_tx_gov_submit-proposal.md) - İlk depozitoyla birlikte bir teklif gönderin 
        - [cancel-software-upgrade \[flags\]](axelard_tx_gov_submit-proposal_cancel-software-upgrade.md) - Mevcut yazılım yükseltme teklifini iptal edin 
        - [community-pool-spend \[teklif-file\]](axelard_tx_gov_submit-proposal_community -pool-spend.md) - Bir topluluk havuzu harcama teklifi gönderin 
        - [ibc-upgrade \[name\] \[height\] \[path/to/upgraded_client_state.json\] \[flags\]](axelard_tx_gov_submit-proposal_ibc -upgrade.md) - Bir IBC yükseltme teklifi gönderin
        - [param-change \[teklif-dosyası\]](axelard_tx_gov_submit-proposal_param-change.md) - Bir parametre değişikliği teklifi gönderin 
        - [yazılım-upgrade \[name\] (--upgrade-height \[height\]) (--upgrade-info \[info\]) \[flags\]](axelard_tx_gov_submit-proposal_software-upgrade.md) - Bir yazılım yükseltme teklifi gönderin 
        - [update-client \[subject-client-id\] \[substitute -client-id\] \[flags\]](axelard_tx_gov_submit-proposal_update-client.md) - Bir güncelleme IBC müşteri teklifi gönderin 
      - [vote \[teklif-id\] \[option\]](axelard_tx_gov_vote.md) - Etkin bir teklif için oy verin, seçenekler: evet/hayır/hayır_with_veto/abstain 
      - [ağırlıklı oy \[teklif kimliği\] \[ağırlıklı seçenekler\]](axelard_tx_gov_weighted-vote.md) - Etkin bir teklif için oy verin, seçenekler : evet/hayır/hayır_with_veto/çekimser
    - [ibc](axelard_tx_ibc.md) - IBC işlem alt komutları 
      - [kanal](axelard_tx_ibc_channel.md) - IBC kanalı işlem alt komutları 
      - [client](axelard_tx_ibc_client.md) - IBC istemci işlem alt komutları 
        - [oluştur \[yol/to/ client_state.json\] \[path/to/consensus_state.json\]](axelard_tx_ibc_client_create.md) - yeni IBC istemcisi oluşturun 
        - [misbehaviour \[path/to/misbehaviour.json\]](axelard_tx_ibc_client_misbehaviour.md) - gönder istemci hatalı davranışı 
        - [güncelleme \[client-id\] \[path/to/header.json\]](axelard_tx_ibc_client_update.md) - mevcut istemciyi bir başlıkla güncelleyin
        - [yükseltme \[client-identifier\] \[path/to/client_state.json\] \[path/to/consensus_state.json\] \[upgrade-client-proof\] \[upgrade-consensus-state-proof \]](axelard_tx_ibc_client_upgrade.md) - bir IBC istemcisini yükseltin 
    - [ibc-transfer](axelard_tx_ibc-transfer.md) - IBC takas edilebilir token transfer işlemi alt komutları 
      - [transfer \[src-port\] \[src-channel\] \[alıcı\] \[amount\]](axelard_tx_ibc-transfer_transfer.md) - IBC aracılığıyla değiştirilebilir bir jeton aktarın 
    - [multisign \[file\] \[name\] \[\[signature\]...\] ](axelard_tx_multisign.md) - Çevrimdışı oluşturulan işlemler için multisig imzaları oluşturun 
    - [multisign-batch \[file\] \[name\] \[\[signature-file\]...\]](axelard_tx_multisign-batch.md ) - Toplu imzalardan toplu olarak çoklu imza işlemlerini birleştirin
    - [nexus](axelard_tx_nexus.md) - nexus işlemleri alt komutları 
      - [activate-chain \[chain\]...](axelard_tx_nexus_activate-chain.md) - verilen zincirleri etkinleştirin 
      - [deactivate-chain \[chain\]. ..](axelard_tx_nexus_deactivate-chain.md) - verilen zincirleri devre dışı bırakın 
      - [deregister-chain-maintainer \[chains\]](axelard_tx_nexus_deregister-chain-maintainer.md) - verilen zincirler için bir doğrulayıcının kaydını zincir bakıcısı olarak silin 
      - [register-chain-maintainer \[chains\]](axelard_tx_nexus_register-chain-maintainer.md) - verilen zincirler için bir doğrulayıcıyı zincir bakımcısı olarak 
    kaydedin - [izin](axelard_tx_permission.md) - izin işlemleri alt komutları
      - [deregister-denetleyici \[controller\]](axelard_tx_permission_deregister-controller.md) - Denetleyici hesabının kaydını sil 
      - [register-controller \[controller\]](axelard_tx_permission_register-controller.md) - Denetleyici hesabını kaydedin 
      - [update-governance- anahtar \[eşik\] \[\[pubKey\]...\]](axelard_tx_permission_update-governance-key.md) - axelar ağı için multisig yönetişim anahtarını güncelleyin 
    - [sign \[file\]](axelard_tx_sign.md ) - Çevrimdışı olarak oluşturulan bir işlemi 
    imzalayın - [sign-batch \[file\]](axelard_tx_sign-batch.md) - İşlem toplu dosyalarını imzalayın 
    - [slashing](axelard_tx_slashing.md) - Slashing işlem alt komutları 
      - [unjail](axelard_tx_slashing_unjail. md) - daha önce kapalı kalma süresi nedeniyle hapse atılan doğrulayıcıyı hapisten çıkarın
    - [anlık görüntü](axelard_tx_snapshot.md) - anlık görüntü işlemleri alt komutları 
      - [deactivate-proxy](axelard_tx_snapshot_deactivate-proxy.md) - Gönderenin proxy hesabını devre dışı bırakın 
      - [register-proxy \[proxy address\]](axelardshot_tx_reg_snapshot .md) - Onun yerine işlemleri yayınlamak üzere belirli bir doğrulayıcı anapara için bir proxy hesabı kaydedin 
      - [gönderme belirteçleri \[miktar\] \[adres 1\] ... \[adres n\]](axelard_tx_snapshot_send-tokens. md) - Belirlenen adreslere belirteç miktarını gönderir 
    - [stake](axelard_tx_staking.md) - Stake işlemi alt komutları 
      - [create-validator](axelard_tx_staking_create-validator.md) - bir kendi kendine yetkilendirme ile başlatılan yeni doğrulayıcı oluşturun o
      - [temsilci \[validator-addr\] \[amount\]](axelard_tx_staking_delegate.md) - Sıvı belirteçlerini bir doğrulayıcıya 
      devredin - [edit-validator](axelard_tx_staking_edit-validator.md) - mevcut bir doğrulayıcı hesabını düzenleyin 
      - [yeniden yetki veren \[src-validator-addr\] \[dst-validator-addr\] \[amount\]](axelard_tx_staking_redelegate.md) - Bir doğrulayıcıdan diğerine likit olmayan belirteçleri yeniden 
      devredin - [unbond \[validator-addr\] \[ miktar\]](axelard_tx_staking_unbond.md) - Bir doğrulayıcıdan hisselerin bağlantısını kaldırın 
    - [tss](axelard_tx_tss.md) - tss işlemleri alt komutları 
      - [register-external-keys \[chain\]](axelard_tx_tss_register-mdsternal) - - Verilen zincir için harici anahtarları kaydedin
      - [döndür \[chain\] \[role\] \[keyID\]](axelard_tx_tss_rotate.md) - Verilen zinciri eski anahtardan verilen anahtara döndürün 
      - [start-keygen](axelard_tx_tss_start-keygen.md) - Anahtar oluşturma protokolünü başlatın 
    - [imzaları doğrulayın \[dosya\]](axelard_tx_validate-signatures.md) - işlem imzalarını doğrulayın 
    - [vesting](axelard_tx_vesting.md) - İşlem alt komutlarını 
      hak etme - [yaratma-hesabı \[adres oluştur] \] \[amount\] \[end_time\]](axelard_tx_vesting_create-vesting-account.md) - Jeton tahsisi ile finanse edilen yeni bir hakediş hesabı oluşturun. 
    - [vesting](axelard_tx_vesting.md) - İşlem alt komutlarını hak etme
      - [create-vesting-account \[to_address\] \[amount\] \[end_time\]](axelard_tx_vesting_create-vesting-account.md) - Jeton tahsisi ile finanse edilen yeni bir hakediş hesabı oluşturun. 
  - [unsafe-reset-all](axelard_unsafe-reset-all.md) - Blok zinciri veritabanını sıfırlar, adres defteri dosyalarını kaldırır ve data/priv_validator_state.json'u başlangıç ​​durumuna sıfırlar 
  - [vald-start](axelard_vald-start. md) - 
  - [validate-genesis \[file\]](axelard_validate-genesis.md) - genesis dosyasını varsayılan konumda veya bir arg olarak geçirilen konumda doğrular 
  - [sürüm](axelard_version.md) - Yazdır uygulama ikili sürüm bilgisi
