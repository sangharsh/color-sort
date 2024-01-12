"use client"

import { GetLevelRequest, PourRequest, ResetRequest, UndoRequest, NextLevelRequest } from '/gen/game_pb.js';
import { ColorSortApiClient } from '/gen/game_grpc_web_pb.js';

const API_URL = 'http://localhost:8080/api';
const METADATA = { 'colorsort-userid': 'abc123' };
const SERVICE = new ColorSortApiClient(API_URL);

export function GetLevel(callback) {
    const request = new GetLevelRequest();

    const processResponse = (err, response) => {
        if (err) {
            console.log("err:", err, "response: ", response);
            return;
        }
        callback(response);
    };

    SERVICE.getLevel(request, METADATA, processResponse);
}

export function Pour(src, dst, callback) {
    const req = new PourRequest();
    req.setSrc(src);
    req.setDst(dst);

    const processResponse = (err, response) => {
        if (err) {
            console.log("err:", err, "response: ", response);
            return;
        }
        callback(response.getLevel());
    };
    SERVICE.pour(req, METADATA, processResponse);
}

export function Reset(callback) {
    const req = new ResetRequest();
    const processResponse = (err, response) => {
        if (err) {
            console.log("err:", err, "response: ", response);
            return;
        }
        callback(response);
    };

    SERVICE.reset(req, METADATA, processResponse);
}

export function Undo(callback) {
    const req = new UndoRequest();
    const processResponse = (err, response) => {
        if (err) {
            console.log("err:", err, "response: ", response);
            return;
        }
        callback(response);
    };

    SERVICE.undo(req, METADATA, processResponse);
}

export function NextLevel(callback) {
    const req = new NextLevelRequest();
    const processResponse = (err, response) => {
        if (err) {
            console.log("err:", err, "response: ", response);
            return;
        }
        callback(response);
    };

    SERVICE.nextLevel(req, METADATA, processResponse);
}