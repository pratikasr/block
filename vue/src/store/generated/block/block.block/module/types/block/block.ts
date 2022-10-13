/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "block.block";

export interface Asset {
  id: number;
  name: string;
  denom: string;
  decimals: number;
  is_on_chain: boolean;
  is_oracle_price_required: boolean;
}

const baseAsset: object = {
  id: 0,
  name: "",
  denom: "",
  decimals: 0,
  is_on_chain: false,
  is_oracle_price_required: false,
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
    if (message.decimals !== 0) {
      writer.uint32(32).int64(message.decimals);
    }
    if (message.is_on_chain === true) {
      writer.uint32(40).bool(message.is_on_chain);
    }
    if (message.is_oracle_price_required === true) {
      writer.uint32(48).bool(message.is_oracle_price_required);
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
          message.decimals = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.is_on_chain = reader.bool();
          break;
        case 6:
          message.is_oracle_price_required = reader.bool();
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
    if (object.decimals !== undefined && object.decimals !== null) {
      message.decimals = Number(object.decimals);
    } else {
      message.decimals = 0;
    }
    if (object.is_on_chain !== undefined && object.is_on_chain !== null) {
      message.is_on_chain = Boolean(object.is_on_chain);
    } else {
      message.is_on_chain = false;
    }
    if (
      object.is_oracle_price_required !== undefined &&
      object.is_oracle_price_required !== null
    ) {
      message.is_oracle_price_required = Boolean(
        object.is_oracle_price_required
      );
    } else {
      message.is_oracle_price_required = false;
    }
    return message;
  },

  toJSON(message: Asset): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.denom !== undefined && (obj.denom = message.denom);
    message.decimals !== undefined && (obj.decimals = message.decimals);
    message.is_on_chain !== undefined &&
      (obj.is_on_chain = message.is_on_chain);
    message.is_oracle_price_required !== undefined &&
      (obj.is_oracle_price_required = message.is_oracle_price_required);
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
    if (object.decimals !== undefined && object.decimals !== null) {
      message.decimals = object.decimals;
    } else {
      message.decimals = 0;
    }
    if (object.is_on_chain !== undefined && object.is_on_chain !== null) {
      message.is_on_chain = object.is_on_chain;
    } else {
      message.is_on_chain = false;
    }
    if (
      object.is_oracle_price_required !== undefined &&
      object.is_oracle_price_required !== null
    ) {
      message.is_oracle_price_required = object.is_oracle_price_required;
    } else {
      message.is_oracle_price_required = false;
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
