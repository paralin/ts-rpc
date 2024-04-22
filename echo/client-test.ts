import { Client, ERR_RPC_ABORT } from '../srpc/index.js'
import { EchoMsg } from './echo_pb.js'
import { EchoerClientImpl } from './echo_srpc.pb.js'
import { pushable } from 'it-pushable'
import { buildRpcStreamOpenStream } from '../rpcstream/rpcstream.js'
import { PartialMessage } from '@bufbuild/protobuf'

export async function runClientTest(client: Client) {
  const demoServiceClient = new EchoerClientImpl(client)

  console.log('Calling Echo: unary call...')
  let result = await demoServiceClient.echo({
    body: 'Hello world!',
  })
  console.log('success: output', result.body)

  // observable for client requests
  const clientRequestStream = pushable<PartialMessage<EchoMsg>>({
    objectMode: true,
  })
  clientRequestStream.push({ body: 'Hello world from streaming request.' })
  clientRequestStream.end()

  console.log('Calling EchoClientStream: client -> server...')
  result = await demoServiceClient.echoClientStream(clientRequestStream)
  console.log('success: output', result.body)

  console.log('Calling EchoServerStream: server -> client...')
  const serverStream = demoServiceClient.echoServerStream({
    body: 'Hello world from server to client streaming request.',
  })
  for await (const msg of serverStream) {
    console.log('server: output', msg.body)
  }
}

// runAbortControllerTest tests aborting a RPC call.
export async function runAbortControllerTest(client: Client) {
  const demoServiceClient = new EchoerClientImpl(client)

  console.log('Testing EchoClientStream with AbortController...')
  let errorReturned = false

  const testRpc = async (rpc: (signal: AbortSignal) => Promise<void>) => {
    const clientAbort = new AbortController()
    new Promise((resolve) => setTimeout(resolve, 1000)).then(() => {
      clientAbort.abort()
    })
    try {
      await rpc(clientAbort.signal)
    } catch (err) {
      const errMsg = (err as Error).message
      errorReturned = true
      if (errMsg !== ERR_RPC_ABORT) {
        throw new Error('unexpected error: ' + errMsg)
      }
    }
    if (!errorReturned) {
      throw new Error('expected aborted rpc to throw error')
    }
  }

  await testRpc(async (signal) => {
    const clientNoopStream = pushable<EchoMsg>({ objectMode: true })
    await demoServiceClient.echoClientStream(clientNoopStream, signal)
  })

  await testRpc(async (signal) => {
    const stream = demoServiceClient.echoServerStream({ body: 'test' }, signal)
    const msgs = []
    try {
      for await (const msg of stream) {
        msgs.push(msg)
      }
    } catch (err) {
      if (msgs.length < 3) {
        throw new Error('expected at least three messages before error')
      }
      throw err
    }
  })
}

// runRpcStreamTest tests a RPCStream.
export async function runRpcStreamTest(client: Client) {
  console.log('Calling RpcStream to open a RPC stream client...')
  const service = new EchoerClientImpl(client)
  const openStreamFn = buildRpcStreamOpenStream(
    'test',
    service.rpcStream.bind(service),
  )
  const proxiedClient = new Client(openStreamFn)
  const proxiedService = new EchoerClientImpl(proxiedClient)

  console.log('Calling Echo via RPC stream...')
  const resp = await proxiedService.echo({ body: 'hello world via proxy' })
  console.log('rpc stream test: succeeded: response: ' + resp.body)

  console.log('Running client test over RPC stream...')
  await runClientTest(proxiedClient)
}
