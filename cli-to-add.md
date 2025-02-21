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

# lnd-cli

```
Usage:
  lnd [OPTIONS]

Application Options:
  -V, --version                                              Display version information and exit
      --lnddir=                                              The base directory that contains lnd's data, logs, configuration
                                                             file, etc. This option overwrites all other directory options.
                                                             (default: /home/kio/.lnd)
  -C, --configfile=                                          Path to configuration file (default: /home/kio/.lnd/lnd.conf)
  -b, --datadir=                                             The directory to store lnd's data within (default:
                                                             /home/kio/.lnd/data)
      --sync-freelist                                        Whether the databases used within lnd should sync their freelist
                                                             to disk. This is disabled by default resulting in improved memory
                                                             performance during operation, but with an increase in startup
                                                             time.
      --tlscertpath=                                         Path to write the TLS certificate for lnd's RPC and REST services
                                                             (default: /home/kio/.lnd/tls.cert)
      --tlskeypath=                                          Path to write the TLS private key for lnd's RPC and REST services
                                                             (default: /home/kio/.lnd/tls.key)
      --tlsextraip=                                          Adds an extra ip to the generated certificate
      --tlsextradomain=                                      Adds an extra domain to the generated certificate
      --tlsautorefresh                                       Re-generate TLS certificate and key if the IPs or domains are
                                                             changed
      --tlsdisableautofill                                   Do not include the interface IPs or the system hostname in TLS
                                                             certificate, use first --tlsextradomain as Common Name instead,
                                                             if set
      --tlscertduration=                                     The duration for which the auto-generated TLS certificate will be
                                                             valid for (default: 10080h0m0s)
      --tlsencryptkey                                        Automatically encrypts the TLS private key and generates
                                                             ephemeral TLS key pairs when the wallet is locked or not
                                                             initialized
      --no-macaroons                                         Disable macaroon authentication, can only be used if server is
                                                             not listening on a public interface.
      --adminmacaroonpath=                                   Path to write the admin macaroon for lnd's RPC and REST services
                                                             if it doesn't exist
      --readonlymacaroonpath=                                Path to write the read-only macaroon for lnd's RPC and REST
                                                             services if it doesn't exist
      --invoicemacaroonpath=                                 Path to the invoice-only macaroon for lnd's RPC and REST services
                                                             if it doesn't exist
      --logdir=                                              Directory to log output. (default: /home/kio/.lnd/logs)
      --maxlogfiles=                                         Maximum logfiles to keep (0 for no rotation) (default: 3)
      --maxlogfilesize=                                      Maximum logfile size in MB (default: 10)
      --acceptortimeout=                                     Time after which an RPCAcceptor will time out and return false if
                                                             it hasn't yet received a response (default: 15s)
      --letsencryptdir=                                      The directory to store Let's Encrypt certificates within
                                                             (default: /home/kio/.lnd/letsencrypt)
      --letsencryptlisten=                                   The IP:port on which lnd will listen for Let's Encrypt
                                                             challenges. Let's Encrypt will always try to contact on port 80.
                                                             Often non-root processes are not allowed to bind to ports lower
                                                             than 1024. This configuration option allows a different port to
                                                             be used, but must be used in combination with port forwarding
                                                             from port 80. This configuration can also be used to specify
                                                             another IP address to listen on, for example an IPv6 address.
                                                             (default: :80)
      --letsencryptdomain=                                   Request a Let's Encrypt certificate for this domain. Note that
                                                             the certificate is only requested and stored when the first rpc
                                                             connection comes in.
      --rpclisten=                                           Add an interface/port/socket to listen for RPC connections
      --restlisten=                                          Add an interface/port/socket to listen for REST connections
      --listen=                                              Add an interface/port to listen for peer connections
      --externalip=                                          Add an ip:port to the list of local addresses we claim to listen
                                                             on to peers. If a port is not specified, the default (9735) will
                                                             be used regardless of other parameters
      --externalhosts=                                       Add a hostname:port that should be periodically resolved to
                                                             announce IPs for. If a port is not specified, the default (9735)
                                                             will be used.
      --restcors=                                            Add an ip:port/hostname to allow cross origin access from. To
                                                             allow all origins, set as "*".
      --nolisten                                             Disable listening for incoming peer connections
      --norest                                               Disable REST API
      --no-rest-tls                                          Disable TLS for REST connections
      --ws-ping-interval=                                    The ping interval for REST based WebSocket connections, set to 0
                                                             to disable sending ping messages from the server side (default:
                                                             30s)
      --ws-pong-wait=                                        The time we wait for a pong response message on REST based
                                                             WebSocket connections before the connection is closed as inactive
                                                             (default: 5s)
      --nat                                                  Toggle NAT traversal support (using either UPnP or NAT-PMP) to
                                                             automatically advertise your external IP address to the network
                                                             -- NOTE this does not support devices behind multiple NATs
      --addpeer=                                             Specify peers to connect to first
      --minbackoff=                                          Shortest backoff when reconnecting to persistent peers. Valid
                                                             time units are {s, m, h}. (default: 1s)
      --maxbackoff=                                          Longest backoff when reconnecting to persistent peers. Valid time
                                                             units are {s, m, h}. (default: 1h0m0s)
      --connectiontimeout=                                   The timeout value for network connections. Valid time units are
                                                             {ms, s, m, h}. (default: 2m0s)
  -d, --debuglevel=                                          Logging level for all subsystems {trace, debug, info, warn,
                                                             error, critical} -- You may also specify
                                                             <global-level>,<subsystem>=<level>,<subsystem2>=<level>,... to
                                                             set the log level for individual subsystems -- Use show to list
                                                             available subsystems (default: info)
      --cpuprofile=                                          Write CPU profile to the specified file
      --profile=                                             Enable HTTP profiling on either a port or host:port
      --blockingprofile=                                     Used to enable a blocking profile to be served on the profiling
                                                             port. This takes a value from 0 to 1, with 1 including every
                                                             blocking event, and 0 including no events.
      --mutexprofile=                                        Used to Enable a mutex profile to be served on the profiling
                                                             port. This takes a value from 0 to 1, with 1 including every
                                                             mutex event, and 0 including no events.
      --unsafe-replay                                        Causes a link to replay the adds on its commitment txn after
                                                             starting up, this enables testing of the sphinx replay logic.
      --maxpendingchannels=                                  The maximum number of incoming pending channels permitted per
                                                             peer. (default: 1)
      --backupfilepath=                                      The target location of the channel backup file
      --blockcachesize=                                      The maximum capacity of the block cache (default: 20971520)
      --nobootstrap                                          If true, then automatic network bootstrapping will not be
                                                             attempted.
      --noseedbackup                                         If true, NO SEED WILL BE EXPOSED -- EVER, AND THE WALLET WILL BE
                                                             ENCRYPTED USING THE DEFAULT PASSPHRASE. THIS FLAG IS ONLY FOR
                                                             TESTING AND SHOULD NEVER BE USED ON MAINNET.
      --wallet-unlock-password-file=                         The full path to a file (or pipe/device) that contains the
                                                             password for unlocking the wallet; if set, no unlocking through
                                                             RPC is possible and lnd will exit if no wallet exists or the
                                                             password is incorrect; if wallet-unlock-allow-create is also set
                                                             then lnd will ignore this flag if no wallet exists and allow a
                                                             wallet to be created through RPC.
      --wallet-unlock-allow-create                           Don't fail with an error if wallet-unlock-password-file is set
                                                             but no wallet exists yet.
      --reset-wallet-transactions                            Removes all transaction history from the on-chain wallet on
                                                             startup, forcing a full chain rescan starting at the wallet's
                                                             birthday. Implements the same functionality as btcwallet's
                                                             dropwtxmgr command. Should be set to false after successful
                                                             execution to avoid rescanning on every restart of lnd.
      --coin-selection-strategy=[largest|random]             The strategy to use for selecting coins for wallet transactions.
                                                             (default: largest)
      --payments-expiration-grace-period=                    A period to wait before force closing channels with outgoing
                                                             htlcs that have timed-out and are a result of this node initiated
                                                             payments.
      --trickledelay=                                        Time in milliseconds between each release of announcements to the
                                                             network (default: 90000)
      --chan-enable-timeout=                                 The duration that a peer connection must be stable before
                                                             attempting to send a channel update to re-enable or cancel a
                                                             pending disables of the peer's channels on the network. (default:
                                                             19m0s)
      --chan-disable-timeout=                                The duration that must elapse after first detecting that an
                                                             already active channel is actually inactive and sending channel
                                                             update disabling it to the network. The pending disable can be
                                                             canceled if the peer reconnects and becomes stable for
                                                             chan-enable-timeout before the disable update is sent. (default:
                                                             20m0s)
      --chan-status-sample-interval=                         The polling interval between attempts to detect if an active
                                                             channel has become inactive due to its peer going offline.
                                                             (default: 1m0s)
      --height-hint-cache-query-disable                      Disable queries from the height-hint cache to try to recover
                                                             channels stuck in the pending close state. Disabling height hint
                                                             queries may cause longer chain rescans, resulting in a
                                                             performance hit. Unset this after channels are unstuck so you can
                                                             get better performance again.
      --alias=                                               The node alias. Used as a moniker by peers and intelligence
                                                             services
      --color=                                               The color of the node in hex format (i.e. '#3399FF'). Used to
                                                             customize node appearance in intelligence services (default:
                                                             #3399FF)
      --minchansize=                                         The smallest channel size (in satoshis) that we should accept.
                                                             Incoming channels smaller than this will be rejected (default:
                                                             20000)
      --maxchansize=                                         The largest channel size (in satoshis) that we should accept.
                                                             Incoming channels larger than this will be rejected
      --coop-close-target-confs=                             The target number of blocks that a cooperative channel close
                                                             transaction should confirm in. This is used to estimate the fee
                                                             to use as the lower bound during fee negotiation for the channel
                                                             closure. (default: 6)
      --channel-commit-interval=                             The maximum time that is allowed to pass between receiving a
                                                             channel state update and signing the next commitment. Setting
                                                             this to a longer duration allows for more efficient channel
                                                             operations at the cost of latency. (default: 50ms)
      --pending-commit-interval=                             The maximum time that is allowed to pass while waiting for the
                                                             remote party to revoke a locally initiated commitment state.
                                                             Setting this to a longer duration if a slow response is expected
                                                             from the remote party or large number of payments are attempted
                                                             at the same time. (default: 1m0s)
      --channel-commit-batch-size=                           The maximum number of channel state updates that is accumulated
                                                             before signing a new commitment. (default: 10)
      --keep-failed-payment-attempts                         Keeps persistent record of all failed payment attempts for
                                                             successfully settled payments.
      --store-final-htlc-resolutions                         Persistently store the final resolution of incoming htlcs.
      --default-remote-max-htlcs=                            The default max_htlc applied when opening or accepting channels.
                                                             This value limits the number of concurrent HTLCs that the remote
                                                             party can add to the commitment. The maximum possible value is
                                                             483. (default: 483)
      --numgraphsyncpeers=                                   The number of peers that we should receive new graph updates
                                                             from. This option can be tuned to save bandwidth for light
                                                             clients or routing nodes. (default: 3)
      --historicalsyncinterval=                              The polling interval between historical graph sync attempts. Each
                                                             historical graph sync attempt ensures we reconcile with the
                                                             remote peer's graph from the genesis block. (default: 1h0m0s)
      --ignore-historical-gossip-filters                     If true, will not reply with historical data that matches the
                                                             range specified by a remote peer's gossip_timestamp_filter. Doing
                                                             so will result in lower memory and bandwidth requirements.
      --rejectpush                                           If true, lnd will not accept channel opening requests with
                                                             non-zero push amounts. This should prevent accidental pushes to
                                                             merchant nodes.
      --rejecthtlc                                           If true, lnd will not forward any HTLCs that are meant as onward
                                                             payments. This option will still allow lnd to send HTLCs and
                                                             receive HTLCs but lnd won't be used as a hop.
      --accept-positive-inbound-fees                         If true, lnd will also allow setting positive inbound fees. By
                                                             default, lnd only allows to set negative inbound fees (an inbound
                                                             "discount") to remain backwards compatible with senders whose
                                                             implementations do not yet support inbound fees.
      --requireinterceptor                                   Whether to always intercept HTLCs, even if no stream is attached
      --stagger-initial-reconnect                            If true, will apply a randomized staggering between 0s and 30s
                                                             when reconnecting to persistent peers on startup. The first 10
                                                             reconnections will be attempted instantly, regardless of the
                                                             flag's value
      --max-cltv-expiry=                                     The maximum number of blocks funds could be locked up for when
                                                             forwarding payments. (default: 2016)
      --max-channel-fee-allocation=                          The maximum percentage of total funds that can be allocated to a
                                                             channel's commitment fee. This only applies for the initiator of
                                                             the channel. Valid values are within [0.1, 1]. (default: 0.5)
      --max-commit-fee-rate-anchors=                         The maximum fee rate in sat/vbyte that will be used for
                                                             commitments of channels of the anchors type. Must be large enough
                                                             to ensure transaction propagation (default: 10)
      --dry-run-migration                                    If true, lnd will abort committing a migration if it would
                                                             otherwise have been successful. This leaves the database
                                                             unmodified, and still compatible with the previously active
                                                             version of lnd.
      --enable-upfront-shutdown                              If true, option upfront shutdown script will be enabled. If peers
                                                             that we open channels with support this feature, we will
                                                             automatically set the script to which cooperative closes should
                                                             be paid out to on channel open. This offers the partial
                                                             protection of a channel peer disconnecting from us if cooperative
                                                             close is attempted with a different script.
      --accept-keysend                                       If true, spontaneous payments through keysend will be accepted.
                                                             [experimental]
      --accept-amp                                           If true, spontaneous payments via AMP will be accepted.
      --keysend-hold-time=                                   If non-zero, keysend payments are accepted but not immediately
                                                             settled. If the payment isn't settled manually after the
                                                             specified time, it is canceled automatically. [experimental]
      --gc-canceled-invoices-on-startup                      If true, we'll attempt to garbage collect canceled invoices upon
                                                             start.
      --gc-canceled-invoices-on-the-fly                      If true, we'll delete newly canceled invoices on the fly.
      --dust-threshold=                                      Sets the dust sum threshold in satoshis for a channel after which
                                                             dust HTLC's will be failed. (default: 500000)
      --allow-circular-route                                 If true, our node will allow htlc forwards that arrive and depart
                                                             on the same channel.
      --http-header-timeout=                                 The maximum duration that the server will wait before timing out
                                                             reading the headers of an HTTP request. (default: 5s)

Bitcoin:
      --bitcoin.chaindir=                                    The directory to store the chain's data within.
      --bitcoin.node=[btcd|bitcoind|neutrino|nochainbackend] The blockchain interface to use. (default: btcd)
      --bitcoin.mainnet                                      Use the main network
      --bitcoin.testnet                                      Use the test network
      --bitcoin.simnet                                       Use the simulation test network
      --bitcoin.regtest                                      Use the regression test network
      --bitcoin.signet                                       Use the signet test network
      --bitcoin.signetchallenge=                             Connect to a custom signet network defined by this challenge
                                                             instead of using the global default signet test network -- Can be
                                                             specified multiple times
      --bitcoin.signetseednode=                              Specify a seed node for the signet network instead of using the
                                                             global default signet network seed nodes
      --bitcoin.defaultchanconfs=                            The default number of confirmations a channel must have before
                                                             it's considered open. If this is not set, we will scale the value
                                                             according to the channel size.
      --bitcoin.defaultremotedelay=                          The default number of blocks we will require our channel
                                                             counterparty to wait before accessing its funds in case of
                                                             unilateral close. If this is not set, we will scale the value
                                                             according to the channel size.
      --bitcoin.maxlocaldelay=                               The maximum blocks we will allow our funds to be timelocked
                                                             before accessing its funds in case of unilateral close. If a peer
                                                             proposes a value greater than this, we will reject the channel.
                                                             (default: 2016)
      --bitcoin.minhtlc=                                     The smallest HTLC we are willing to accept on our channels, in
                                                             millisatoshi (default: 1)
      --bitcoin.minhtlcout=                                  The smallest HTLC we are willing to send out on our channels, in
                                                             millisatoshi (default: 1000)
      --bitcoin.basefee=                                     The base fee in millisatoshi we will charge for forwarding
                                                             payments on our channels (default: 1000)
      --bitcoin.feerate=                                     The fee rate used when forwarding payments on our channels. The
                                                             total fee charged is basefee + (amount * feerate / 1000000),
                                                             where amount is the forwarded amount. (default: 1)
      --bitcoin.timelockdelta=                               The CLTV delta we will subtract from a forwarded HTLC's timelock
                                                             value (default: 80)
      --bitcoin.dnsseed=                                     The seed DNS server(s) to use for initial peer discovery. Must be
                                                             specified as a '<primary_dns>[,<soa_primary_dns>]' tuple where
                                                             the SOA address is needed for DNS resolution through Tor but is
                                                             optional for clearnet users. Multiple tuples can be specified,
                                                             will overwrite the default seed servers.

btcd:
      --btcd.dir=                                            The base directory that contains the node's data, logs,
                                                             configuration file, etc. (default: /home/kio/.btcd)
      --btcd.rpchost=                                        The daemon's rpc listening address. If a port is omitted, then
                                                             the default port for the selected chain parameters will be used.
                                                             (default: localhost)
      --btcd.rpcuser=                                        Username for RPC connections
      --btcd.rpcpass=                                        Password for RPC connections
      --btcd.rpccert=                                        File containing the daemon's certificate file (default:
                                                             /home/kio/.btcd/rpc.cert)
      --btcd.rawrpccert=                                     The raw bytes of the daemon's PEM-encoded certificate chain which
                                                             will be used to authenticate the RPC connection.

bitcoind:
      --bitcoind.dir=                                        The base directory that contains the node's data, logs,
                                                             configuration file, etc. (default: /home/kio/.bitcoin)
      --bitcoind.config=                                     Configuration filepath. If not set, will default to the default
                                                             filename under 'dir'.
      --bitcoind.rpccookie=                                  Authentication cookie file for RPC connections. If not set, will
                                                             default to .cookie under 'dir'.
      --bitcoind.rpchost=                                    The daemon's rpc listening address. If a port is omitted, then
                                                             the default port for the selected chain parameters will be used.
                                                             (default: localhost)
      --bitcoind.rpcuser=                                    Username for RPC connections
      --bitcoind.rpcpass=                                    Password for RPC connections
      --bitcoind.zmqpubrawblock=                             The address listening for ZMQ connections to deliver raw block
                                                             notifications
      --bitcoind.zmqpubrawtx=                                The address listening for ZMQ connections to deliver raw
                                                             transaction notifications
      --bitcoind.zmqreaddeadline=                            The read deadline for reading ZMQ messages from both the block
                                                             and tx subscriptions (default: 5s)
      --bitcoind.estimatemode=                               The fee estimate mode. Must be either ECONOMICAL or CONSERVATIVE.
                                                             (default: CONSERVATIVE)
      --bitcoind.pruned-node-max-peers=                      The maximum number of peers lnd will choose from the backend node
                                                             to retrieve pruned blocks from. This only applies to pruned
                                                             nodes. (default: 4)
      --bitcoind.rpcpolling                                  Poll the bitcoind RPC interface for block and transaction
                                                             notifications instead of using the ZMQ interface
      --bitcoind.blockpollinginterval=                       The interval that will be used to poll bitcoind for new blocks.
                                                             Only used if rpcpolling is true.
      --bitcoind.txpollinginterval=                          The interval that will be used to poll bitcoind for new tx. Only
                                                             used if rpcpolling is true.

neutrino:
  -a, --neutrino.addpeer=                                    Add a peer to connect with at startup
      --neutrino.connect=                                    Connect only to the specified peers at startup
      --neutrino.maxpeers=                                   Max number of inbound and outbound peers
      --neutrino.banduration=                                How long to ban misbehaving peers.  Valid time units are {s, m,
                                                             h}.  Minimum 1 second
      --neutrino.banthreshold=                               Maximum allowed ban score before disconnecting and banning
                                                             misbehaving peers.
      --neutrino.assertfilterheader=                         Optional filter header in height:hash format to assert the state
                                                             of neutrino's filter header chain on startup. If the assertion
                                                             does not hold, then the filter header chain will be re-synced
                                                             from the genesis block.
      --neutrino.useragentname=                              Used to help identify ourselves to other bitcoin peers (default:
                                                             neutrino)
      --neutrino.useragentversion=                           Used to help identify ourselves to other bitcoin peers (default:
                                                             0.12.0-beta)
      --neutrino.validatechannels                            Validate every channel in the graph during sync by downloading
                                                             the containing block. This is the inverse of
                                                             routing.assumechanvalid, meaning that for Neutrino the validation
                                                             is turned off by default for massively increased graph sync
                                                             performance. This speedup comes at the risk of using an
                                                             unvalidated view of the network for routing. Overwrites the value
                                                             of routing.assumechanvalid if Neutrino is used. (default: false)
      --neutrino.broadcasttimeout=                           The amount of time to wait before giving up on a transaction
                                                             broadcast attempt.
      --neutrino.persistfilters                              Whether compact filters fetched from the P2P network should be
                                                             persisted to disk.

Autopilot:
      --autopilot.active                                     If the autopilot agent should be active or not.
      --autopilot.heuristic=                                 Heuristic to activate, and the weight to give it during scoring.
                                                             (default: {top_centrality:1})
      --autopilot.maxchannels=                               The maximum number of channels that should be created (default: 5)
      --autopilot.allocation=                                The percentage of total funds that should be committed to
                                                             automatic channel establishment (default: 0.6)
      --autopilot.minchansize=                               The smallest channel that the autopilot agent should create
                                                             (default: 20000)
      --autopilot.maxchansize=                               The largest channel that the autopilot agent should create
                                                             (default: 16777215)
      --autopilot.private                                    Whether the channels created by the autopilot agent should be
                                                             private or not. Private channels won't be announced to the
                                                             network.
      --autopilot.minconfs=                                  The minimum number of confirmations each of your inputs in
                                                             funding transactions created by the autopilot agent must have.
                                                             (default: 1)
      --autopilot.conftarget=                                The confirmation target (in blocks) for channels opened by
                                                             autopilot. (default: 3)

Tor:
      --tor.active                                           Allow outbound and inbound connections to be routed through Tor
      --tor.socks=                                           The host:port that Tor's exposed SOCKS5 proxy is listening on
                                                             (default: localhost:9050)
      --tor.dns=                                             The DNS server as host:port that Tor will use for SRV queries -
                                                             NOTE must have TCP resolution enabled (default:
                                                             soa.nodes.lightning.directory:53)
      --tor.streamisolation                                  Enable Tor stream isolation by randomizing user credentials for
                                                             each connection.
      --tor.skip-proxy-for-clearnet-targets                  Allow the node to establish direct connections to services not
                                                             running behind Tor.
      --tor.control=                                         The host:port that Tor is listening on for Tor control
                                                             connections (default: localhost:9051)
      --tor.targetipaddress=                                 IP address that Tor should use as the target of the hidden service
      --tor.password=                                        The password used to arrive at the HashedControlPassword for the
                                                             control port. If provided, the HASHEDPASSWORD authentication
                                                             method will be used instead of the SAFECOOKIE one.
      --tor.v2                                               Automatically set up a v2 onion service to listen for inbound
                                                             connections
      --tor.v3                                               Automatically set up a v3 onion service to listen for inbound
                                                             connections
      --tor.privatekeypath=                                  The path to the private key of the onion service being created
      --tor.encryptkey                                       Encrypts the Tor private key file on disk
      --tor.watchtowerkeypath=                               The path to the private key of the watchtower onion service being
                                                             created

routerrpc:
      --routerrpc.estimator=[apriori|bimodal]                Probability estimator used for pathfinding. (default: apriori)
      --routerrpc.minrtprob=                                 Minimum required route success probability to attempt the payment
                                                             (default: 0.01)
      --routerrpc.attemptcost=                               The fixed (virtual) cost in sats of a failed payment attempt
                                                             (default: 100)
      --routerrpc.attemptcostppm=                            The proportional (virtual) cost in sats of a failed payment
                                                             attempt expressed in parts per million of the total payment
                                                             amount (default: 1000)
      --routerrpc.maxmchistory=                              the maximum number of payment results that are held on disk by
                                                             mission control (default: 1000)
      --routerrpc.mcflushinterval=                           the timer interval to use to flush mission control state to the
                                                             DB (default: 1s)
      --routerrpc.fee-estimation-timeout=                    the maximum time to wait for routing fees to be estimated by
                                                             payment probes (default: 1m0s)
      --routerrpc.usestatusinitiated                         If true, the router will send Payment_INITIATED for new payments,
                                                             otherwise Payment_In_FLIGHT will be sent for compatibility
                                                             concerns.
      --routerrpc.routermacaroonpath=                        Path to the router macaroon

apriori:
      --routerrpc.apriori.hopprob=                           Assumed success probability of a hop in a route when no other
                                                             information is available. (default: 0.6)
      --routerrpc.apriori.weight=                            Weight of the a priori probability in success probability
                                                             estimation. Valid values are in [0, 1]. (default: 0.5)
      --routerrpc.apriori.penaltyhalflife=                   Defines the duration after which a penalized node or channel is
                                                             back at 50% probability (default: 1h0m0s)
      --routerrpc.apriori.capacityfraction=                  Defines the fraction of channels' capacities that is considered
                                                             liquid. Valid values are in [0.75, 1]. (default: 0.9999)

bimodal:
      --routerrpc.bimodal.scale=                             Defines the unbalancedness assumed for the network, the amount
                                                             defined in msat. (default: 300000000)
      --routerrpc.bimodal.nodeweight=                        Defines how strongly non-routed channels should be taken into
                                                             account for probability estimation. Valid values are in [0, 1].
                                                             (default: 0.2)
      --routerrpc.bimodal.decaytime=                         Describes the information decay of knowledge about previous
                                                             successes and failures in channels. (default: 168h0m0s)

fee:
      --fee.url=                                             Optional URL for external fee estimation. If no URL is specified,
                                                             the method for fee estimation will depend on the chosen backend
                                                             and network. Must be set for neutrino on mainnet.
      --fee.min-update-timeout=                              The minimum interval in which fees will be updated from the
                                                             specified fee URL. (default: 5m0s)
      --fee.max-update-timeout=                              The maximum interval in which fees will be updated from the
                                                             specified fee URL. (default: 20m0s)

invoices:
      --invoices.holdexpirydelta=                            The number of blocks before a hold invoice's htlc expires that
                                                             the invoice should be canceled to prevent a force close. Force
                                                             closes will not be prevented if this value is not greater than
                                                             DefaultIncomingBroadcastDelta. (default: 12)

routing:
      --routing.strictgraphpruning                           If true, then the graph will be pruned more aggressively for
                                                             zombies. In practice this means that edges with a single stale
                                                             edge will be considered a zombie.

gossip:
      --gossip.pinned-syncers=                               A set of peers that should always remain in an active sync state,
                                                             which can be used to closely synchronize the routing tables of
                                                             two nodes. The value should be a hex-encoded pubkey, the flag can
                                                             be specified multiple times to add multiple peers. Connected
                                                             peers matching this pubkey will remain active for the duration of
                                                             the connection and not count towards the NumActiveSyncer count.
      --gossip.max-channel-update-burst=                     The maximum number of updates for a specific channel and
                                                             direction that lnd will accept over the channel update interval.
                                                             (default: 10)
      --gossip.channel-update-interval=                      The interval used to determine how often lnd should allow a burst
                                                             of new updates for a specific channel and direction. (default:
                                                             1m0s)
      --gossip.sub-batch-delay=                              The duration to wait before sending the next announcement batch
                                                             if there are multiple. Use a small value if there are a lot
                                                             announcements and they need to be broadcast quickly. (default: 5s)

workers:
      --workers.read=                                        Maximum number of concurrent read pool workers. This number
                                                             should be proportional to the number of peers. (default: 100)
      --workers.write=                                       Maximum number of concurrent write pool workers. This number
                                                             should be proportional to the number of CPUs on the host.
                                                             (default: 8)
      --workers.sig=                                         Maximum number of concurrent sig pool workers. This number should
                                                             be proportional to the number of CPUs on the host. (default: 8)

caches:
      --caches.reject-cache-size=                            Maximum number of entries contained in the reject cache, which is
                                                             used to speed up filtering of new channel announcements and
                                                             channel updates from peers. Each entry requires 25 bytes.
                                                             (default: 50000)
      --caches.channel-cache-size=                           Maximum number of entries contained in the channel cache, which
                                                             is used to reduce memory allocations from gossip queries from
                                                             peers. Each entry requires roughly 2Kb. (default: 20000)
      --caches.rpc-graph-cache-duration=                     The period of time expressed as a duration (1s, 1m, 1h, etc) that
                                                             the RPC response to DescribeGraph should be cached for.

wtclient:
      --wtclient.active                                      Whether the daemon should use private watchtowers to back up
                                                             revoked channel states.
      --wtclient.sweep-fee-rate=                             Specifies the fee rate in sat/byte to be used when constructing
                                                             justice transactions sent to the watchtower. (default: 10)
      --wtclient.session-close-range=                        The range over which to choose a random number of blocks to wait
                                                             after the last channel of a session is closed before sending the
                                                             DeleteSession message to the tower server. Set to 1 for no delay.
                                                             (default: 288)
      --wtclient.max-tasks-in-mem-queue=                     The maximum number of updates that should be queued in memory
                                                             before overflowing to disk. (default: 2000)
      --wtclient.max-updates=                                The maximum number of updates to be backed up in a single
                                                             session. (default: 1024)

watchtower:
      --watchtower.active                                    If the watchtower should be active or not
      --watchtower.towerdir=                                 Directory of the watchtower.db (default:
                                                             /home/kio/.lnd/data/watchtower)
      --watchtower.listen=                                   Add interfaces/ports to listen for peer connections
      --watchtower.externalip=                               Add interfaces/ports where the watchtower can accept peer
                                                             connections
      --watchtower.readtimeout=                              Duration the watchtower server will wait for messages to be
                                                             received before hanging up on clients (default: 15s)
      --watchtower.writetimeout=                             Duration the watchtower server will wait for messages to be
                                                             written before hanging up on client connections (default: 15s)

protocol:
      --protocol.wumbo-channels                              if set, then lnd will create and accept requests for channels
                                                             larger chan 0.16 BTC
      --protocol.simple-taproot-chans                        if set, then lnd will create and accept requests for channels
                                                             using the simple taproot commitment type
      --protocol.no-anchors                                  disable support for anchor commitments
      --protocol.no-script-enforced-lease                    disable support for script enforced lease commitments
      --protocol.option-scid-alias                           enable support for option_scid_alias channels
      --protocol.zero-conf                                   enable support for zero-conf channels, must have
                                                             option-scid-alias set also
      --protocol.no-any-segwit                               disallow using any segwit witness version as a co-op close address
      --protocol.no-timestamp-query-option                   do not query syncing peers for announcement timestamps and do not
                                                             respond with timestamps if requested
      --protocol.no-route-blinding                           do not forward payments that are a part of a blinded route
      --protocol.custom-message=                             allows the custom message apis to send and report messages with
                                                             the protocol number provided that fall outside of the custom
                                                             message number range.
      --protocol.custom-init=                                custom feature bits — numbers defined in BOLT 9 — to
                                                             advertise in the node's init message
      --protocol.custom-nodeann=                             custom feature bits — numbers defined in BOLT 9 — to
                                                             advertise in the node's announcement message
      --protocol.custom-invoice=                             custom feature bits — numbers defined in BOLT 9 — to
                                                             advertise in the node's invoices

chainbackend:
      --healthcheck.chainbackend.interval=                   How often to run a health check. (default: 1m0s)
      --healthcheck.chainbackend.attempts=                   The number of calls we will make for the check before failing.
                                                             Set this value to 0 to disable a check. (default: 3)
      --healthcheck.chainbackend.timeout=                    The amount of time we allow the health check to take before
                                                             failing due to timeout. (default: 30s)
      --healthcheck.chainbackend.backoff=                    The amount of time to back-off between failed health checks.
                                                             (default: 2m0s)

diskspace:
      --healthcheck.diskspace.diskrequired=                  The minimum ratio of free disk space to total capacity that we
                                                             allow before shutting lnd down safely. (default: 0.1)
      --healthcheck.diskspace.interval=                      How often to run a health check. (default: 12h0m0s)
      --healthcheck.diskspace.attempts=                      The number of calls we will make for the check before failing.
                                                             Set this value to 0 to disable a check.
      --healthcheck.diskspace.timeout=                       The amount of time we allow the health check to take before
                                                             failing due to timeout. (default: 5s)
      --healthcheck.diskspace.backoff=                       The amount of time to back-off between failed health checks.
                                                             (default: 1m0s)

tls:
      --healthcheck.tls.interval=                            How often to run a health check. (default: 1m0s)
      --healthcheck.tls.attempts=                            The number of calls we will make for the check before failing.
                                                             Set this value to 0 to disable a check.
      --healthcheck.tls.timeout=                             The amount of time we allow the health check to take before
                                                             failing due to timeout. (default: 5s)
      --healthcheck.tls.backoff=                             The amount of time to back-off between failed health checks.
                                                             (default: 1m0s)

torconnection:
      --healthcheck.torconnection.interval=                  How often to run a health check. (default: 1m0s)
      --healthcheck.torconnection.attempts=                  The number of calls we will make for the check before failing.
                                                             Set this value to 0 to disable a check.
      --healthcheck.torconnection.timeout=                   The amount of time we allow the health check to take before
                                                             failing due to timeout. (default: 5s)
      --healthcheck.torconnection.backoff=                   The amount of time to back-off between failed health checks.
                                                             (default: 1m0s)

remotesigner:
      --healthcheck.remotesigner.interval=                   How often to run a health check. (default: 1m0s)
      --healthcheck.remotesigner.attempts=                   The number of calls we will make for the check before failing.
                                                             Set this value to 0 to disable a check. (default: 1)
      --healthcheck.remotesigner.timeout=                    The amount of time we allow the health check to take before
                                                             failing due to timeout. (default: 1s)
      --healthcheck.remotesigner.backoff=                    The amount of time to back-off between failed health checks.
                                                             (default: 30s)

db:
      --db.backend=                                          The selected database backend. (default: bolt)
      --db.batch-commit-interval=                            The maximum duration the channel graph batch schedulers will wait
                                                             before attempting to commit a batch of pending updates. This can
                                                             be tradeoff database contenion for commit latency. (default:
                                                             500ms)
      --db.use-native-sql                                    Use native SQL for tables that already support it.
      --db.no-graph-cache                                    Don't use the in-memory graph cache for path finding. Much slower
                                                             but uses less RAM. Can only be used with a bolt database backend.
      --db.prune-revocation                                  Run the optional migration that prunes the revocation logs to
                                                             save disk space.
      --db.no-rev-log-amt-data                               If set, the to-local and to-remote output amounts of revoked
                                                             commitment transactions will not be stored in the revocation log.
                                                             Note that once this data is lost, a watchtower client will not be
                                                             able to back up the revoked state.

etcd:
      --db.etcd.embedded                                     Use embedded etcd instance instead of the external one. Note: use
                                                             for testing only.
      --db.etcd.embedded_client_port=                        Client port to use for the embedded instance. Note: use for
                                                             testing only.
      --db.etcd.embedded_peer_port=                          Peer port to use for the embedded instance. Note: use for testing
                                                             only.
      --db.etcd.embedded_log_file=                           Optional log file to use for embedded instance logs. note: use
                                                             for testing only.
      --db.etcd.host=                                        Etcd database host.
      --db.etcd.user=                                        Etcd database user.
      --db.etcd.pass=                                        Password for the database user.
      --db.etcd.namespace=                                   The etcd namespace to use.
      --db.etcd.disabletls                                   Disable TLS for etcd connection. Caution: use for development
                                                             only.
      --db.etcd.cert_file=                                   Path to the TLS certificate for etcd RPC.
      --db.etcd.key_file=                                    Path to the TLS private key for etcd RPC.
      --db.etcd.insecure_skip_verify                         Whether we intend to skip TLS verification
      --db.etcd.collect_stats                                Whether to collect etcd commit stats.
      --db.etcd.max_msg_size=                                The maximum message size in bytes that we may send to etcd.
                                                             (default: 33554432)

bolt:
      --db.bolt.nofreelistsync                               Whether the databases used within lnd should sync their freelist
                                                             to disk. This is set to true by default, meaning we don't sync
                                                             the free-list resulting in improved memory performance during
                                                             operation, but with an increase in startup time.
      --db.bolt.auto-compact                                 Whether the databases used within lnd should automatically be
                                                             compacted on every startup (and if the database has the
                                                             configured minimum age). This is disabled by default because it
                                                             requires additional disk space to be available during the
                                                             compaction that is freed afterwards. In general compaction leads
                                                             to smaller database files.
      --db.bolt.auto-compact-min-age=                        How long ago the last compaction of a database file must be for
                                                             it to be considered for auto compaction again. Can be set to 0 to
                                                             compact on every startup. (default: 168h0m0s)
      --db.bolt.dbtimeout=                                   Specify the timeout value used when opening the database.
                                                             (default: 1m0s)

postgres:
      --db.postgres.dsn=                                     Database connection string.
      --db.postgres.timeout=                                 Database connection timeout. Set to zero to disable.
      --db.postgres.maxconnections=                          The maximum number of open connections to the database. Set to
                                                             zero for unlimited. (default: 50)
      --db.postgres.skipmigrations                           Skip applying migrations on startup.

sqlite:
      --db.sqlite.timeout=                                   The time after which a database query should be timed out.
      --db.sqlite.busytimeout=                               The maximum amount of time to wait for a database connection to
                                                             become available for a query. (default: 5s)
      --db.sqlite.maxconnections=                            The maximum number of open connections to the database. Set to
                                                             zero for unlimited. (default: 2)
      --db.sqlite.pragmaoptions=                             A list of pragma options to set on a database connection. For
                                                             example, 'auto_vacuum=incremental'. Note that the flag must be
                                                             specified multiple times if multiple options are to be set.
      --db.sqlite.skipmigrations                             Skip applying migrations on startup.

cluster:
      --cluster.enable-leader-election                       Enables leader election if set.
      --cluster.leader-elector=[etcd]                        Leader elector to use. Valid values: "etcd". (default: etcd)
      --cluster.etcd-election-prefix=                        Election key prefix when using etcd leader elector. (default:
                                                             /leader/)
      --cluster.id=                                          Identifier for this node inside the cluster (used in leader
                                                             election). Defaults to the hostname. (default: archlinux)
      --cluster.leader-session-ttl=                          The TTL in seconds to use for the leader election session.
                                                             (default: 60)

rpcmiddleware:
      --rpcmiddleware.enable                                 Enable the RPC middleware interceptor functionality.
      --rpcmiddleware.intercepttimeout=                      Time after which a RPC middleware intercept request will time out
                                                             and return an error if it hasn't yet received a response.
                                                             (default: 2s)
      --rpcmiddleware.addmandatory=                          Add the named middleware to the list of mandatory middlewares.
                                                             All RPC requests are blocked/denied if any of the mandatory
                                                             middlewares is not registered. Can be specified multiple times.

remotesigner:
      --remotesigner.enable                                  Use a remote signer for signing any on-chain related transactions
                                                             or messages. Only recommended if local wallet is initialized as
                                                             watch-only. Remote signer must use the same seed/root key as the
                                                             local watch-only wallet but must have private keys.
      --remotesigner.rpchost=                                The remote signer's RPC host:port
      --remotesigner.macaroonpath=                           The macaroon to use for authenticating with the remote signer
      --remotesigner.tlscertpath=                            The TLS certificate to use for establishing the remote signer's
                                                             identity
      --remotesigner.timeout=                                The timeout for connecting to and signing requests with the
                                                             remote signer. Valid time units are {s, m, h}. (default: 5s)
      --remotesigner.migrate-wallet-to-watch-only            If a wallet with private key material already exists, migrate it
                                                             into a watch-only wallet on first startup. WARNING: This cannot
                                                             be undone! Make sure you have backed up your seed before you use
                                                             this flag! All private keys will be purged from the wallet after
                                                             first unlock with this flag!

sweeper:
      --sweeper.maxfeerate=                                  Maximum fee rate in sat/vb that the sweeper is allowed to use
                                                             when sweeping funds, the fee rate derived from budgets are capped
                                                             at this value. Setting this value too low can result in
                                                             transactions not being confirmed in time, causing HTLCs to expire
                                                             hence potentially losing funds. (default: 1000)
      --sweeper.nodeadlineconftarget=                        The conf target to use when sweeping non-time-sensitive outputs.
                                                             This is useful for sweeping outputs that are not time-sensitive,
                                                             and can be swept at a lower fee rate. (default: 1008)
      --sweeper.budget=                                      An optional config group that's used for the automatic sweep fee
                                                             estimation. The Budget config gives options to limits ones fee
                                                             exposure when sweeping unilateral close outputs and the fee rate
                                                             calculated from budgets is capped at sweeper.maxfeerate. Check
                                                             the budget config options for more details.

sweeper.budget:
      --sweeper.budget.tolocal=                              The amount in satoshis to allocate as the budget to pay fees when
                                                             sweeping the to_local output. If set, the budget calculated using
                                                             the ratio (if set) will be capped at this value.
      --sweeper.budget.tolocalratio=                         The ratio of the value in to_local output to allocate as the
                                                             budget to pay fees when sweeping it. (default: 0.5)
      --sweeper.budget.anchorcpfp=                           The amount in satoshis to allocate as the budget to pay fees when
                                                             CPFPing a force close tx using the anchor output. If set, the
                                                             budget calculated using the ratio (if set) will be capped at this
                                                             value.
      --sweeper.budget.anchorcpfpratio=                      The ratio of a special value to allocate as the budget to pay
                                                             fees when CPFPing a force close tx using the anchor output. The
                                                             special value is the sum of all time-sensitive HTLCs on this
                                                             commitment subtracted by their budgets. (default: 0.5)
      --sweeper.budget.deadlinehtlc=                         The amount in satoshis to allocate as the budget to pay fees when
                                                             sweeping a time-sensitive (first-level) HTLC. If set, the budget
                                                             calculated using the ratio (if set) will be capped at this value.
      --sweeper.budget.deadlinehtlcratio=                    The ratio of the value in a time-sensitive (first-level) HTLC to
                                                             allocate as the budget to pay fees when sweeping it. (default:
                                                             0.5)
      --sweeper.budget.nodeadlinehtlc=                       The amount in satoshis to allocate as the budget to pay fees when
                                                             sweeping a non-time-sensitive (second-level) HTLC. If set, the
                                                             budget calculated using the ratio (if set) will be capped at this
                                                             value.
      --sweeper.budget.nodeadlinehtlcratio=                  The ratio of the value in a non-time-sensitive (second-level)
                                                             HTLC to allocate as the budget to pay fees when sweeping it.
                                                             (default: 0.5)

htlcswitch:
      --htlcswitch.mailboxdeliverytimeout=                   The timeout value when delivering HTLCs to a channel link.
                                                             Setting this value too small will result in local payment
                                                             failures if large number of payments are sent over a short
                                                             period. (default: 1m0s)

grpc:
      --grpc.server-ping-time=                               How long the server waits on a gRPC stream with no activity
                                                             before pinging the client. (default: 1m0s)
      --grpc.server-ping-timeout=                            How long the server waits for the response from the client for
                                                             the keepalive ping response. (default: 20s)
      --grpc.client-ping-min-wait=                           The minimum amount of time the client should wait before sending
                                                             a keepalive ping. (default: 5s)
      --grpc.client-allow-ping-without-stream                If true, the server allows keepalive pings from the client even
                                                             when there are no active gRPC streams. This might be useful to
                                                             keep the underlying HTTP/2 connection open for future requests.

Help Options:
  -h, --help                                                 Show this help message
```
