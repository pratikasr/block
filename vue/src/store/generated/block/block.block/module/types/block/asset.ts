/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "block.block";

export interface Asset {
  id: number;
  name: string;
  denom: string;
  decimal: string;
  price: string;
  appId: string;
  ibcStatus: string;
}

const baseAsset: object = {
  id: 0,
  name: "",
  denom: "",
  decimal: "",
  price: "",
  appId: "",
  ibcStatus: "",
};

export const Asset = {
  encode(message: Asset, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
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

  decode(input: Reader | Uint8Array, length?: number): Asset {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAsset } as Asset;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
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

  fromJSON(object: any): Asset {
    const message = { ...baseAsset } as Asset;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
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

  toJSON(message: Asset): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.denom !== undefined && (obj.denom = message.denom);
    message.decimal !== undefined && (obj.decimal = message.decimal);
    message.price !== undefined && (obj.price = message.price);
    message.appId !== undefined && (obj.appId = message.appId);
    message.ibcStatus !== undefined && (obj.ibcStatus = message.ibcStatus);
    return obj;
  },

  fromPartial(object: DeepPartial<Asset>): Asset {
    const message = { ...baseAsset } as Asset;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
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

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
