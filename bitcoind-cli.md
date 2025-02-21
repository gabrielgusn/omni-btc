# electrum
```
electrum help
usage: electrum [-h] [--version] [-v VERBOSITY] [-V VERBOSITY_SHORTCUTS] [-D ELECTRUM_PATH] [-P] [--testnet] [--testnet4]
                [--regtest] [--simnet] [--signet] [-o] [--rpcuser RPCUSER] [--rpcpassword RPCPASSWORD] [-w WALLET_PATH]
                [--forgetconfig]
                <command> ...

positional arguments:
  <command>
    gui                 Run GUI (default)
    daemon              Run Daemon
    add_peer
    add_request         Create a payment request, using the first unused address of the wallet
    addtransaction      Add a transaction to the wallet history
    broadcast           Broadcast a transaction to the network
    bumpfee             Bump the fee for an unconfirmed transaction
    changegaplimit      Change the gap limit of the wallet
    clear_invoices      Remove all invoices
    clear_ln_blacklist
    clear_requests      Remove all payment requests
    close_channel
    close_wallet        Close wallet
    commands            List of commands
    convert_currency    Converts the given amount of currency to another using the configured exchange rate source
    convert_xkey        Convert xtype of a master key
    create              Create a new wallet
    createmultisig      Create multisig address
    createnewaddress    Create a new receiving address, beyond the gap limit of the wallet
    decode_invoice
    decrypt             Decrypt a message encrypted with a public key
    delete_invoice      Remove an outgoing payment invoice
    delete_request      Remove an incoming payment request
    deserialize         Deserialize a serialized transaction
    dumpprivkeys        Deprecated
    enable_htlc_settle
    encrypt             Encrypt a message with a public key
    export_channel_backup
    freeze              Freeze address
    freeze_utxo         Freeze a UTXO so that the wallet will not spend it
    get                 Return item from wallet storage
    get_channel_ctx     return the current commitment transaction of a channel
    get_invoice         Returns an invoice (request for outgoing payment)
    get_request         Returns a payment request
    get_tx_status       Returns some information regarding the tx
    get_watchtower_ctn  return the local watchtower's ctn of channel
    getaddressbalance   Return the balance of any address
    getaddresshistory   Return the transaction history of any address
    getaddressunspent   Returns the UTXO list of any address
    getalias            Retrieve alias
    getbalance          Return the balance of your wallet
    getconfig           Return a configuration variable
    getfeerate          Return current fee rate settings and current estimate (in sat/kvByte)
    getinfo             network info
    getmasterprivate    Get master private key
    getmerkle           Get Merkle branch of a transaction included in a block
    getminacceptablegap
                        Returns the minimum value for gap limit that would be sufficient to discover all known addresses in
                        the wallet
    getmpk              Get master public key
    getprivatekeyforpath
                        Get private key corresponding to derivation path (address index)
    getprivatekeys      Get private keys of addresses
    getpubkeys          Return the public keys for a wallet address
    getseed             Get seed phrase
    getservers          Return the list of known servers (candidates for connecting)
    gettransaction      Retrieve a transaction
    getunusedaddress    Returns the first unused address of the wallet, or None if all addresses are used
    help
    import_channel_backup
    importprivkey       Import a private key
    is_synchronized     return wallet synchronization status
    ismine              Check if address is in wallet
    lightning_history   lightning history
    list_channels
    list_invoices       Returns the list of invoices (requests for outgoing payments) saved in the wallet
    list_peers
    list_requests       Returns the list of incoming payment requests saved in the wallet
    list_wallets        List wallets open in daemon
    listaddresses       List wallet addresses
    listcontacts        Show your list of contacts
    listunspent         List unspent outputs
    lnpay
    load_wallet         Load the wallet in memory
    make_seed           Create a seed
    nodeid
    normal_swap         Normal submarine swap: send on-chain BTC, receive on Lightning Note that your funds will be locked
                        for 24h if you do not have enough incoming capacity
    notify              Watch an address
    onchain_history     Wallet onchain history
    open_channel
    password            Change wallet password
    payto               Create a transaction
    paytomany           Create a multi-output transaction
    rebalance_channels  Rebalance channels
    removelocaltx       Remove a 'local' transaction from the wallet, and its dependent transactions
    request_force_close
                        Requests the remote to force close a channel
    reset_liquidity_hints
    restore             Restore a wallet from text
    reverse_swap        Reverse submarine swap: send on Lightning, receive on-chain
    searchcontacts      Search through contacts, return matching entries
    serialize           Create a signed raw transaction from a json tx template
    setconfig           Set a configuration variable
    setfeerate          Set fee rate estimation method and value
    setlabel            Assign a label to an item
    signmessage         Sign a message with a key
    signtransaction     Sign a transaction
    signtransaction_with_privkey
                        Sign a transaction
    stop                Stop daemon
    sweep               Sweep private keys
    unfreeze            Unfreeze address
    unfreeze_utxo       Unfreeze a UTXO so that the wallet might spend it
    validateaddress     Check that an address is valid
    verifymessage       Verify a signature
    version             Return the version of Electrum
    version_info        Return information about dependencies, such as their version and path

options:
  -h, --help            show this help message and exit
  --version             Return the version of Electrum.
  -w, --wallet WALLET_PATH
                        wallet path
  --forgetconfig        Forget config on exit

global options:
  -v VERBOSITY          Set verbosity (log levels)
  -V VERBOSITY_SHORTCUTS
                        Set verbosity (shortcut-filter list)
  -D, --dir ELECTRUM_PATH
                        electrum directory
  -P, --portable        Use local 'electrum_data' directory
  --testnet             Use Testnet
  --testnet4            Use Testnet4
  --regtest             Use Regtest
  --simnet              Use Simnet
  --signet              Use Signet
  -o, --offline         Run offline
  --rpcuser RPCUSER     RPC user
  --rpcpassword RPCPASSWORD
                        RPC password
```

# bitcoin-cli
```
bitcoin-cli --help
Bitcoin Core RPC client version v27.0.0

Usage:  bitcoin-cli [options] <command> [params]  Send command to Bitcoin Core
or:     bitcoin-cli [options] -named <command> [name=value]...  Send command to Bitcoin Core (with named arguments)
or:     bitcoin-cli [options] help                List commands
or:     bitcoin-cli [options] help <command>      Get help for a command

Options:

  -?
       Print this help message and exit

  -addrinfo
       Get the number of addresses known to the node, per network and total,
       after filtering for quality and recency. The total number of
       addresses known to the node may be higher.

  -color=<when>
       Color setting for CLI output (default: auto). Valid values: always, auto
       (add color codes when standard output is connected to a terminal
       and OS is not WIN32), never.

  -conf=<file>
       Specify configuration file. Relative paths will be prefixed by datadir
       location. (default: bitcoin.conf)

  -datadir=<dir>
       Specify data directory

  -generate
       Generate blocks, equivalent to RPC getnewaddress followed by RPC
       generatetoaddress. Optional positional integer arguments are
       number of blocks to generate (default: 1) and maximum iterations
       to try (default: 1000000), equivalent to RPC generatetoaddress
       nblocks and maxtries arguments. Example: bitcoin-cli -generate 4
       1000

  -getinfo
       Get general information from the remote server. Note that unlike
       server-side RPC calls, the output of -getinfo is the result of
       multiple non-atomic requests. Some entries in the output may
       represent results from different states (e.g. wallet balance may
       be as of a different block from the chain state reported)

  -named
       Pass named instead of positional arguments (default: false)

  -netinfo
       Get network peer connection information from the remote server. An
       optional integer argument from 0 to 4 can be passed for different
       peers listings (default: 0). Pass "help" for detailed help
       documentation.

  -rpcclienttimeout=<n>
       Timeout in seconds during HTTP requests, or 0 for no timeout. (default:
       900)

  -rpcconnect=<ip>
       Send commands to node running on <ip> (default: 127.0.0.1)

  -rpccookiefile=<loc>
       Location of the auth cookie. Relative paths will be prefixed by a
       net-specific datadir location. (default: data dir)

  -rpcpassword=<pw>
       Password for JSON-RPC connections

  -rpcport=<port>
       Connect to JSON-RPC on <port> (default: 8332, testnet: 18332, signet:
       38332, regtest: 18443)

  -rpcuser=<user>
       Username for JSON-RPC connections

  -rpcwait
       Wait for RPC server to start

  -rpcwaittimeout=<n>
       Timeout in seconds to wait for the RPC server to start, or 0 for no
       timeout. (default: 0)

  -rpcwallet=<walletname>
       Send RPC for non-default wallet on RPC server (needs to exactly match
       corresponding -wallet option passed to bitcoind). This changes
       the RPC endpoint used, e.g.
       http://127.0.0.1:8332/wallet/<walletname>

  -stdin
       Read extra arguments from standard input, one per line until EOF/Ctrl-D
       (recommended for sensitive information such as passphrases). When
       combined with -stdinrpcpass, the first line from standard input
       is used for the RPC password.

  -stdinrpcpass
       Read RPC password from standard input as a single line. When combined
       with -stdin, the first line from standard input is used for the
       RPC password. When combined with -stdinwalletpassphrase,
       -stdinrpcpass consumes the first line, and -stdinwalletpassphrase
       consumes the second.

  -stdinwalletpassphrase
       Read wallet passphrase from standard input as a single line. When
       combined with -stdin, the first line from standard input is used
       for the wallet passphrase.

  -version
       Print version and exit

Debugging/Testing options:

Chain selection options:

  -chain=<chain>
       Use the chain <chain> (default: main). Allowed values: main, test,
       signet, regtest

  -signet
       Use the signet chain. Equivalent to -chain=signet. Note that the network
       is defined by the -signetchallenge parameter

  -signetchallenge
       Blocks must satisfy the given script to be considered valid (only for
       signet networks; defaults to the global default signet test
       network challenge)

  -signetseednode
       Specify a seed node for the signet network, in the hostname[:port]
       format, e.g. sig.net:1234 (may be used multiple times to specify
       multiple seed nodes; defaults to the global default signet test
       network seed node(s))

  -testnet
       Use the test chain. Equivalent to -chain=test.
3m

```

# bitcoind
```
bitcoind --help
Bitcoin Core version v27.0.0

Usage:  bitcoind [options]                     Start Bitcoin Core

Options:

  -?
       Print this help message and exit

  -alertnotify=<cmd>
       Execute command when an alert is raised (%s in cmd is replaced by
       message)

  -allowignoredconf
       For backwards compatibility, treat an unused bitcoin.conf file in the
       datadir as a warning, not an error.

  -assumevalid=<hex>
       If this block is in the chain assume that it and its ancestors are valid
       and potentially skip their script verification (0 to verify all,
       default:
       000000000000000000026811d149d4d261995ec5b3f64f439a0a10e1a464af9a,
       testnet:
       000000000001323071f38f21ea5aae529ece491eadaccce506a59bcc2d968917,
       signet:
       0000000870f15246ba23c16e370a7ffb1fc8a3dcf8cb4492882ed4b0e3d4cd26)

  -blockfilterindex=<type>
       Maintain an index of compact filters by block (default: 0, values:
       basic). If <type> is not supplied or if <type> = 1, indexes for
       all known types are enabled.

  -blocknotify=<cmd>
       Execute command when the best block changes (%s in cmd is replaced by
       block hash)

  -blockreconstructionextratxn=<n>
       Extra transactions to keep in memory for compact block reconstructions
       (default: 100)

  -blocksdir=<dir>
       Specify directory to hold blocks subdirectory for *.dat files (default:
       <datadir>)

  -blocksonly
       Whether to reject transactions from network peers. Automatic broadcast
       and rebroadcast of any transactions from inbound peers is
       disabled, unless the peer has the 'forcerelay' permission. RPC
       transactions are not affected. (default: 0)

  -coinstatsindex
       Maintain coinstats index used by the gettxoutsetinfo RPC (default: 0)

  -conf=<file>
       Specify path to read-only configuration file. Relative paths will be
       prefixed by datadir location (only useable from command line, not
       configuration file) (default: bitcoin.conf)

  -daemon
       Run in the background as a daemon and accept commands (default: 0)

  -daemonwait
       Wait for initialization to be finished before exiting. This implies
       -daemon (default: 0)

  -datadir=<dir>
       Specify data directory

  -dbcache=<n>
       Maximum database cache size <n> MiB (4 to 16384, default: 450). In
       addition, unused mempool memory is shared for this cache (see
       -maxmempool).

  -debuglogfile=<file>
       Specify location of debug log file (default: debug.log). Relative paths
       will be prefixed by a net-specific datadir location. Pass
       -nodebuglogfile to disable writing the log to a file.

  -includeconf=<file>
       Specify additional configuration file, relative to the -datadir path
       (only useable from configuration file, not command line)

  -loadblock=<file>
       Imports blocks from external file on startup

  -maxmempool=<n>
       Keep the transaction memory pool below <n> megabytes (default: 300)

  -maxorphantx=<n>
       Keep at most <n> unconnectable transactions in memory (default: 100)

  -mempoolexpiry=<n>
       Do not keep transactions in the mempool longer than <n> hours (default:
       336)

  -par=<n>
       Set the number of script verification threads (0 = auto, up to 15, <0 =
       leave that many cores free, default: 0)

  -persistmempool
       Whether to save the mempool on shutdown and load on restart (default: 1)

  -persistmempoolv1
       Whether a mempool.dat file created by -persistmempool or the savemempool
       RPC will be written in the legacy format (version 1) or the
       current format (version 2). This temporary option will be removed
       in the future. (default: 0)

  -pid=<file>
       Specify pid file. Relative paths will be prefixed by a net-specific
       datadir location. (default: bitcoind.pid)

  -prune=<n>
       Reduce storage requirements by enabling pruning (deleting) of old
       blocks. This allows the pruneblockchain RPC to be called to
       delete specific blocks and enables automatic pruning of old
       blocks if a target size in MiB is provided. This mode is
       incompatible with -txindex. Warning: Reverting this setting
       requires re-downloading the entire blockchain. (default: 0 =
       disable pruning blocks, 1 = allow manual pruning via RPC, >=550 =
       automatically prune block files to stay under the specified
       target size in MiB)

  -reindex
       If enabled, wipe chain state and block index, and rebuild them from
       blk*.dat files on disk. Also wipe and rebuild other optional
       indexes that are active. If an assumeutxo snapshot was loaded,
       its chainstate will be wiped as well. The snapshot can then be
       reloaded via RPC.

  -reindex-chainstate
       If enabled, wipe chain state, and rebuild it from blk*.dat files on
       disk. If an assumeutxo snapshot was loaded, its chainstate will
       be wiped as well. The snapshot can then be reloaded via RPC.

  -settings=<file>
       Specify path to dynamic settings data file. Can be disabled with
       -nosettings. File is written at runtime and not meant to be
       edited by users (use bitcoin.conf instead for custom settings).
       Relative paths will be prefixed by datadir location. (default:
       settings.json)

  -shutdownnotify=<cmd>
       Execute command immediately before beginning shutdown. The need for
       shutdown may be urgent, so be careful not to delay it long (if
       the command doesn't require interaction with the server, consider
       having it fork into the background).

  -startupnotify=<cmd>
       Execute command on startup.

  -txindex
       Maintain a full transaction index, used by the getrawtransaction rpc
       call (default: 0)

  -version
       Print version and exit

Connection options:

  -addnode=<ip>
       Add a node to connect to and attempt to keep the connection open (see
       the addnode RPC help for more info). This option can be specified
       multiple times to add multiple nodes; connections are limited to
       8 at a time and are counted separately from the -maxconnections
       limit.

  -asmap=<file>
       Specify asn mapping used for bucketing of the peers (default:
       ip_asn.map). Relative paths will be prefixed by the net-specific
       datadir location.

  -bantime=<n>
       Default duration (in seconds) of manually configured bans (default:
       86400)

  -bind=<addr>[:<port>][=onion]
       Bind to given address and always listen on it (default: 0.0.0.0). Use
       [host]:port notation for IPv6. Append =onion to tag any incoming
       connections to that address and port as incoming Tor connections
       (default: 127.0.0.1:8334=onion, testnet: 127.0.0.1:18334=onion,
       signet: 127.0.0.1:38334=onion, regtest: 127.0.0.1:18445=onion)

  -cjdnsreachable
       If set, then this host is configured for CJDNS (connecting to fc00::/8
       addresses would lead us to the CJDNS network, see doc/cjdns.md)
       (default: 0)

  -connect=<ip>
       Connect only to the specified node; -noconnect disables automatic
       connections (the rules for this peer are the same as for
       -addnode). This option can be specified multiple times to connect
       to multiple nodes.

  -discover
       Discover own IP addresses (default: 1 when listening and no -externalip
       or -proxy)

  -dns
       Allow DNS lookups for -addnode, -seednode and -connect (default: 1)

  -dnsseed
       Query for peer addresses via DNS lookup, if low on addresses (default: 1
       unless -connect used or -maxconnections=0)

  -externalip=<ip>
       Specify your own public address

  -fixedseeds
       Allow fixed seeds if DNS seeds don't provide peers (default: 1)

  -forcednsseed
       Always query for peer addresses via DNS lookup (default: 0)

  -i2pacceptincoming
       Whether to accept inbound I2P connections (default: 1). Ignored if
       -i2psam is not set. Listening for inbound I2P connections is done
       through the SAM proxy, not by binding to a local address and
       port.

  -i2psam=<ip:port>
       I2P SAM proxy to reach I2P peers and accept I2P connections (default:
       none)

  -listen
       Accept connections from outside (default: 1 if no -proxy, -connect or
       -maxconnections=0)

  -listenonion
       Automatically create Tor onion service (default: 1)

  -maxconnections=<n>
       Maintain at most <n> automatic connections to peers (default: 125). This
       limit does not apply to connections manually added via -addnode
       or the addnode RPC, which have a separate limit of 8.

  -maxreceivebuffer=<n>
       Maximum per-connection receive buffer, <n>*1000 bytes (default: 5000)

  -maxsendbuffer=<n>
       Maximum per-connection memory usage for the send buffer, <n>*1000 bytes
       (default: 1000)

  -maxtimeadjustment
       Maximum allowed median peer time offset adjustment. Local perspective of
       time may be influenced by outbound peers forward or backward by
       this amount (default: 4200 seconds).

  -maxuploadtarget=<n>
       Tries to keep outbound traffic under the given target per 24h. Limit
       does not apply to peers with 'download' permission or blocks
       created within past week. 0 = no limit (default: 0M). Optional
       suffix units [k|K|m|M|g|G|t|T] (default: M). Lowercase is 1000
       base while uppercase is 1024 base

  -natpmp
       Use NAT-PMP to map the listening port (default: 0)

  -networkactive
       Enable all P2P network activity (default: 1). Can be changed by the
       setnetworkactive RPC command

  -onion=<ip:port>
       Use separate SOCKS5 proxy to reach peers via Tor onion services, set
       -noonion to disable (default: -proxy)

  -onlynet=<net>
       Make automatic outbound connections only to network <net> (ipv4, ipv6,
       onion, i2p, cjdns). Inbound and manual connections are not
       affected by this option. It can be specified multiple times to
       allow multiple networks.

  -peerblockfilters
       Serve compact block filters to peers per BIP 157 (default: 0)

  -peerbloomfilters
       Support filtering of blocks and transaction with bloom filters (default:
       0)

  -port=<port>
       Listen for connections on <port>. Nodes not using the default ports
       (default: 8333, testnet: 18333, signet: 38333, regtest: 18444)
       are unlikely to get incoming connections. Not relevant for I2P
       (see doc/i2p.md).

  -proxy=<ip:port>
       Connect through SOCKS5 proxy, set -noproxy to disable (default:
       disabled)

  -proxyrandomize
       Randomize credentials for every proxy connection. This enables Tor
       stream isolation (default: 1)

  -seednode=<ip>
       Connect to a node to retrieve peer addresses, and disconnect. This
       option can be specified multiple times to connect to multiple
       nodes.

  -timeout=<n>
       Specify socket connection timeout in milliseconds. If an initial attempt
       to connect is unsuccessful after this amount of time, drop it
       (minimum: 1, default: 5000)

  -torcontrol=<ip>:<port>
       Tor control host and port to use if onion listening enabled (default:
       127.0.0.1:9051). If no port is specified, the default port of
       9051 will be used.

  -torpassword=<pass>
       Tor control port password (default: empty)

  -upnp
       Use UPnP to map the listening port (default: 1 when listening and no
       -proxy)

  -v2transport
       Support v2 transport (default: 1)

  -whitebind=<[permissions@]addr>
       Bind to the given address and add permission flags to the peers
       connecting to it. Use [host]:port notation for IPv6. Allowed
       permissions: bloomfilter (allow requesting BIP37 filtered blocks
       and transactions), noban (do not ban for misbehavior; implies
       download), forcerelay (relay transactions that are already in the
       mempool; implies relay), relay (relay even in -blocksonly mode,
       and unlimited transaction announcements), mempool (allow
       requesting BIP35 mempool contents), download (allow getheaders
       during IBD, no disconnect after maxuploadtarget limit), addr
       (responses to GETADDR avoid hitting the cache and contain random
       records with the most up-to-date info). Specify multiple
       permissions separated by commas (default:
       download,noban,mempool,relay). Can be specified multiple times.

  -whitelist=<[permissions@]IP address or network>
       Add permission flags to the peers connecting from the given IP address
       (e.g. 1.2.3.4) or CIDR-notated network (e.g. 1.2.3.0/24). Uses
       the same permissions as -whitebind. Can be specified multiple
       times.

Wallet options:

  -addresstype
       What type of addresses to use ("legacy", "p2sh-segwit", "bech32", or
       "bech32m", default: "bech32")

  -avoidpartialspends
       Group outputs by address, selecting many (possibly all) or none, instead
       of selecting on a per-output basis. Privacy is improved as
       addresses are mostly swept with fewer transactions and outputs
       are aggregated in clean change addresses. It may result in higher
       fees due to less optimal coin selection caused by this added
       limitation and possibly a larger-than-necessary number of inputs
       being used. Always enabled for wallets with "avoid_reuse"
       enabled, otherwise default: 0.

  -changetype
       What type of change to use ("legacy", "p2sh-segwit", "bech32", or
       "bech32m"). Default is "legacy" when -addresstype=legacy, else it
       is an implementation detail.

  -consolidatefeerate=<amt>
       The maximum feerate (in BTC/kvB) at which transaction building may use
       more inputs than strictly necessary so that the wallet's UTXO
       pool can be reduced (default: 0.0001).

  -disablewallet
       Do not load the wallet and disable wallet RPC calls

  -discardfee=<amt>
       The fee rate (in BTC/kvB) that indicates your tolerance for discarding
       change by adding it to the fee (default: 0.0001). Note: An output
       is discarded if it is dust at this rate, but we will always
       discard up to the dust relay fee and a discard fee above that is
       limited by the fee estimate for the longest target

  -fallbackfee=<amt>
       A fee rate (in BTC/kvB) that will be used when fee estimation has
       insufficient data. 0 to entirely disable the fallbackfee feature.
       (default: 0.00)

  -keypool=<n>
       Set key pool size to <n> (default: 1000). Warning: Smaller sizes may
       increase the risk of losing funds when restoring from an old
       backup, if none of the addresses in the original keypool have
       been used.

  -maxapsfee=<n>
       Spend up to this amount in additional (absolute) fees (in BTC) if it
       allows the use of partial spend avoidance (default: 0.00)

  -mintxfee=<amt>
       Fee rates (in BTC/kvB) smaller than this are considered zero fee for
       transaction creation (default: 0.00001)

  -paytxfee=<amt>
       Fee rate (in BTC/kvB) to add to transactions you send (default: 0.00)

  -signer=<cmd>
       External signing tool, see doc/external-signer.md

  -spendzeroconfchange
       Spend unconfirmed change when sending transactions (default: 1)

  -txconfirmtarget=<n>
       If paytxfee is not set, include enough fee so transactions begin
       confirmation on average within n blocks (default: 6)

  -wallet=<path>
       Specify wallet path to load at startup. Can be used multiple times to
       load multiple wallets. Path is to a directory containing wallet
       data and log files. If the path is not absolute, it is
       interpreted relative to <walletdir>. This only loads existing
       wallets and does not create new ones. For backwards compatibility
       this also accepts names of existing top-level data files in
       <walletdir>.

  -walletbroadcast
       Make the wallet broadcast transactions (default: 1)

  -walletdir=<dir>
       Specify directory to hold wallets (default: <datadir>/wallets if it
       exists, otherwise <datadir>)

  -walletnotify=<cmd>
       Execute command when a wallet transaction changes. %s in cmd is replaced
       by TxID, %w is replaced by wallet name, %b is replaced by the
       hash of the block including the transaction (set to 'unconfirmed'
       if the transaction is not included) and %h is replaced by the
       block height (-1 if not included). %w is not currently
       implemented on windows. On systems where %w is supported, it
       should NOT be quoted because this would break shell escaping used
       to invoke the command.

  -walletrbf
       Send transactions with full-RBF opt-in enabled (RPC only, default: 1)

ZeroMQ notification options:

  -zmqpubhashblock=<address>
       Enable publish hash block in <address>

  -zmqpubhashblockhwm=<n>
       Set publish hash block outbound message high water mark (default: 1000)

  -zmqpubhashtx=<address>
       Enable publish hash transaction in <address>

  -zmqpubhashtxhwm=<n>
       Set publish hash transaction outbound message high water mark (default:
       1000)

  -zmqpubrawblock=<address>
       Enable publish raw block in <address>

  -zmqpubrawblockhwm=<n>
       Set publish raw block outbound message high water mark (default: 1000)

  -zmqpubrawtx=<address>
       Enable publish raw transaction in <address>

  -zmqpubrawtxhwm=<n>
       Set publish raw transaction outbound message high water mark (default:
       1000)

  -zmqpubsequence=<address>
       Enable publish hash block and tx sequence in <address>

  -zmqpubsequencehwm=<n>
       Set publish hash sequence message high water mark (default: 1000)

Debugging/Testing options:

  -debug=<category>
       Output debug and trace logging (default: -nodebug, supplying <category>
       is optional). If <category> is not supplied or if <category> = 1,
       output all debug and trace logging. <category> can be: addrman,
       bench, blockstorage, cmpctblock, coindb, estimatefee, http, i2p,
       ipc, leveldb, libevent, mempool, mempoolrej, net, proxy, prune,
       qt, rand, reindex, rpc, scan, selectcoins, tor, txpackages,
       txreconciliation, util, validation, walletdb, zmq. This option
       can be specified multiple times to output multiple categories.

  -debugexclude=<category>
       Exclude debug and trace logging for a category. Can be used in
       conjunction with -debug=1 to output debug and trace logging for
       all categories except the specified category. This option can be
       specified multiple times to exclude multiple categories.

  -help-debug
       Print help message with debugging options and exit

  -logips
       Include IP addresses in debug output (default: 0)

  -loglevelalways
       Always prepend a category and level (default: 0)

  -logsourcelocations
       Prepend debug output with name of the originating source location
       (source file, line number and function name) (default: 0)

  -logthreadnames
       Prepend debug output with name of the originating thread (only available
       on platforms supporting thread_local) (default: 0)

  -logtimestamps
       Prepend debug output with timestamp (default: 1)

  -maxtxfee=<amt>
       Maximum total fees (in BTC) to use in a single wallet transaction;
       setting this too low may abort large transactions (default: 0.10)

  -printtoconsole
       Send trace/debug info to console (default: 1 when no -daemon. To disable
       logging to file, set -nodebuglogfile)

  -shrinkdebugfile
       Shrink debug.log file on client startup (default: 1 when no -debug)

  -uacomment=<cmt>
       Append comment to the user agent string

Chain selection options:

  -chain=<chain>
       Use the chain <chain> (default: main). Allowed values: main, test,
       signet, regtest

  -signet
       Use the signet chain. Equivalent to -chain=signet. Note that the network
       is defined by the -signetchallenge parameter

  -signetchallenge
       Blocks must satisfy the given script to be considered valid (only for
       signet networks; defaults to the global default signet test
       network challenge)

  -signetseednode
       Specify a seed node for the signet network, in the hostname[:port]
       format, e.g. sig.net:1234 (may be used multiple times to specify
       multiple seed nodes; defaults to the global default signet test
       network seed node(s))

  -testnet
       Use the test chain. Equivalent to -chain=test.

Node relay options:

  -bytespersigop
       Equivalent bytes per sigop in transactions for relay and mining
       (default: 20)

  -datacarrier
       Relay and mine data carrier transactions (default: 1)

  -datacarriersize
       Relay and mine transactions whose data-carrying raw scriptPubKey is of
       this size or less (default: 83)

  -mempoolfullrbf
       Accept transaction replace-by-fee without requiring replaceability
       signaling (default: 0)

  -minrelaytxfee=<amt>
       Fees (in BTC/kvB) smaller than this are considered zero fee for
       relaying, mining and transaction creation (default: 0.00001)

  -permitbaremultisig
       Relay non-P2SH multisig (default: 1)

  -whitelistforcerelay
       Add 'forcerelay' permission to whitelisted inbound peers with default
       permissions. This will relay transactions even if the
       transactions were already in the mempool. (default: 0)

  -whitelistrelay
       Add 'relay' permission to whitelisted inbound peers with default
       permissions. This will accept relayed transactions even when not
       relaying transactions (default: 1)

Block creation options:

  -blockmaxweight=<n>
       Set maximum BIP141 block weight (default: 3996000)

  -blockmintxfee=<amt>
       Set lowest fee rate (in BTC/kvB) for transactions to be included in
       block creation. (default: 0.00001)

RPC server options:

  -rest
       Accept public REST requests (default: 0)

  -rpcallowip=<ip>
       Allow JSON-RPC connections from specified source. Valid values for <ip>
       are a single IP (e.g. 1.2.3.4), a network/netmask (e.g.
       1.2.3.4/255.255.255.0), a network/CIDR (e.g. 1.2.3.4/24), all
       ipv4 (0.0.0.0/0), or all ipv6 (::/0). This option can be
       specified multiple times

  -rpcauth=<userpw>
       Username and HMAC-SHA-256 hashed password for JSON-RPC connections. The
       field <userpw> comes in the format: <USERNAME>:<SALT>$<HASH>. A
       canonical python script is included in share/rpcauth. The client
       then connects normally using the
       rpcuser=<USERNAME>/rpcpassword=<PASSWORD> pair of arguments. This
       option can be specified multiple times

  -rpcbind=<addr>[:port]
       Bind to given address to listen for JSON-RPC connections. Do not expose
       the RPC server to untrusted networks such as the public internet!
       This option is ignored unless -rpcallowip is also passed. Port is
       optional and overrides -rpcport. Use [host]:port notation for
       IPv6. This option can be specified multiple times (default:
       127.0.0.1 and ::1 i.e., localhost)

  -rpccookiefile=<loc>
       Location of the auth cookie. Relative paths will be prefixed by a
       net-specific datadir location. (default: data dir)

  -rpcpassword=<pw>
       Password for JSON-RPC connections

  -rpcport=<port>
       Listen for JSON-RPC connections on <port> (default: 8332, testnet:
       18332, signet: 38332, regtest: 18443)

  -rpcthreads=<n>
       Set the number of threads to service RPC calls (default: 4)

  -rpcuser=<user>
       Username for JSON-RPC connections

  -rpcwhitelist=<whitelist>
       Set a whitelist to filter incoming RPC calls for a specific user. The
       field <whitelist> comes in the format: <USERNAME>:<rpc 1>,<rpc
       2>,...,<rpc n>. If multiple whitelists are set for a given user,
       they are set-intersected. See -rpcwhitelistdefault documentation
       for information on default whitelist behavior.

  -rpcwhitelistdefault
       Sets default behavior for rpc whitelisting. Unless rpcwhitelistdefault
       is set to 0, if any -rpcwhitelist is set, the rpc server acts as
       if all rpc users are subject to empty-unless-otherwise-specified
       whitelists. If rpcwhitelistdefault is set to 1 and no
       -rpcwhitelist is set, rpc server acts as if all rpc users are
       subject to empty whitelists.

  -server
       Accept command line and JSON-RPC commands
```
