'use client'

import './styles.css';

import { useState } from 'react';
import { NewLevelPlayRequest } from '/gen/game_pb.js';
import { ColorSortApiClient } from '/gen/game_grpc_web_pb.js';

export default function Page() {
    const [game, setGame] = useState({});
    return (
        <Game game={game} setGame={setGame} />
    )
}

function Game({ game, setGame }) {
    const renderedTubes = [];
    const [selected, setSelected] = useState(-1);

    function handleTubeSelection(tubeIndex) {
        if (selected == tubeIndex) {
            setSelected(-1);
            return;
        }
        if (selected != -1 && selected != tubeIndex) {
            pour(game.getTubesList()[selected], game.getTubesList()[tubeIndex]);
            setGame(game);
            setSelected(-1);
            if (hasWon(game.getTubesList())) {
                console.log("Won!!!");
            }
            return;
        }
        setSelected(tubeIndex);
    }

    if (!game || !game.array) {
        grpcCall(setGame);
    }
    if (game.array) {
        game.getTubesList().forEach((tube, index) => {
            renderedTubes.push(<Tube tube={tube} key={index} tubeIndex={index} selected={selected == index} handleTubeSelection={handleTubeSelection} />);
        });
    }

    return (
        <div className="container">
            <h1>Level {game.level}</h1>
            {renderedTubes}
        </div>
    )
}


function pour(src, dst) {
    if (dst.getColorsList().length == dst.getSize()) {
        console.log("dst full");
        return;
    }
    if (dst.getColorsList().length != 0 && src.getColorsList()[src.getColorsList().length - 1] != dst.getColorsList()[dst.getColorsList().length - 1]) {
        console.log("colors non matching:", src.getColorsList()[src.getColorsList().length - 1], dst.getColorsList()[src.getColorsList().length - 1]);
        return;
    }
    console.log("colors, ", src.getColorsList()[src.getColorsList().length - 1], dst.getColorsList()[dst.getColorsList().length - 1]);
    while (dst.getColorsList().length == 0 || (dst.getColorsList().length != dst.getSize() && src.getColorsList()[src.getColorsList().length - 1] == dst.getColorsList()[dst.getColorsList().length - 1])) {
        dst.getColorsList().push(src.getColorsList().pop());
    }
}

function hasWon(tubes) {
    return tubes.every((tube) => {
        if (tube.getColorsList().length == 0)
            return true;
        if (tube.getColorsList().length != tube.getSize())
            return false;
        let color = tube.getColorsList()[0];
        return tube.getColorsList().every(e => e == color);
    })
}

function Tube({ tube, tubeIndex, selected, handleTubeSelection }) {
    const levels = [];
    for (var i = tube.getSize() - 1; i >= 0; i--) {
        let color = i < tube.getColorsList().length ? tube.getColorsList()[i] : '';
        levels.push(<TubeColor color={color} key={i} />);
    }

    return (
        <div className={"testtube" + (selected ? ' selected' : '')} onClick={e => handleTubeSelection(tubeIndex)}>
            {levels}
        </div>
    );
}

function TubeColor({ color }) {
    return (<div className={`liquid liquid-${color}`}></div>)
}

function grpcCall(callback) {
    var service = new ColorSortApiClient('http://localhost:8080');
    const request = new NewLevelPlayRequest();
    request.setId(1);

    const processResponse = (err, response) => {
        console.log("err:", err, "response: ", response);
        callback(response.getCurrentstate());
    };

    service.newLevel(request, {}, processResponse);

}
