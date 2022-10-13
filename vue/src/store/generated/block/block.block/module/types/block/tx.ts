/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "block.block";

export interface MsgAddAsset {
  creator: string;
  name: string;
  denom: string;
  decimal: string;
  price: string;
  appId: string;
  ibcStatus: string;
}

export interface MsgAddAssetResponse {}

const baseMsgAddAsset: object = {
  creator: "",
  name: "",
  denom: "",
  decimal: "",
  price: "",
  appId: "",
  ibcStatus: "",
};

export const MsgAddAsset = {
  encode(message: MsgAddAsset, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.denom !== "") {
      writer.uint32(26).string(message.denom);
    }
    if (message.decimal !== "") {
      writer.uint32(34).string(message.decimal);
    }
    if (message.price !== "") {
      writer.uint32(42).string(message.price);
    }
    if (message.appId !== "") {
      writer.uint32(50).string(message.appId);
    }
    if (message.ibcStatus !== "") {
      writer.uint32(58).string(message.ibcStatus);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddAsset {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddAsset } as MsgAddAsset;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.denom = reader.string();
          break;
        case 4:
          message.decimal = reader.string();
          break;
        case 5:
          message.price = reader.string();
          break;
        case 6:
          message.appId = reader.string();
          break;
        case 7:
          message.ibcStatus = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddAsset {
    const message = { ...baseMsgAddAsset } as MsgAddAsset;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    if (object.decimal !== undefined && object.decimal !== null) {
      message.decimal = String(object.decimal);
    } else {
      message.decimal = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = String(object.price);
    } else {
      message.price = "";
    }
    if (object.appId !== undefined && object.appId !== null) {
      message.appId = String(object.appId);
    } else {
      message.appId = "";
    }
    if (object.ibcStatus !== undefined && object.ibcStatus !== null) {
      message.ibcStatus = String(object.ibcStatus);
    } else {
      message.ibcStatus = "";
    }
    return message;
  },

  toJSON(message: MsgAddAsset): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.denom !== undefined && (obj.denom = message.denom);
    message.decimal !== undefined && (obj.decimal = message.decimal);
    message.price !== undefined && (obj.price = message.price);
    message.appId !== undefined && (obj.appId = message.appId);
    message.ibcStatus !== undefined && (obj.ibcStatus = message.ibcStatus);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddAsset>): MsgAddAsset {
    const message = { ...baseMsgAddAsset } as MsgAddAsset;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    if (object.decimal !== undefined && object.decimal !== null) {
      message.decimal = object.decimal;
    } else {
      message.decimal = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = "";
    }
    if (object.appId !== undefined && object.appId !== null) {
      message.appId = object.appId;
    } else {
      message.appId = "";
    }
    if (object.ibcStatus !== undefined && object.ibcStatus !== null) {
      message.ibcStatus = object.ibcStatus;
    } else {
      message.ibcStatus = "";
    }
    return message;
  },
};

const baseMsgAddAssetResponse: object = {};

export const MsgAddAssetResponse = {
  encode(_: MsgAddAssetResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddAssetResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddAssetResponse } as MsgAddAssetResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAddAssetResponse {
    const message = { ...baseMsgAddAssetResponse } as MsgAddAssetResponse;
    return message;
  },

  toJSON(_: MsgAddAssetResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgAddAssetResponse>): MsgAddAssetResponse {
    const message = { ...baseMsgAddAssetResponse } as MsgAddAssetResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AddAsset(request: MsgAddAsset): Promise<MsgAddAssetResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  AddAsset(request: MsgAddAsset): Promise<MsgAddAssetResponse> {
    const data = MsgAddAsset.encode(request).finish();
    const promise = this.rpc.request("block.block.Msg", "AddAsset", data);
    return promise.then((data) => MsgAddAssetResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
