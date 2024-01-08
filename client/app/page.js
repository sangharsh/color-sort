"use client"

import './styles.css';

import { useState } from 'react';
import { NewLevel, Pour, Reset, Undo } from './service';

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
        if (selected == -1) { // Src select
            setSelected(tubeIndex);
        } else if (selected == tubeIndex) { // Same selection -> unselect
            setSelected(-1);
        } else if (selected != -1) { // Dst select -> Pour
            Pour(selected, tubeIndex, setGame);
            setSelected(-1);
        }
    }

    if (!game || !game.array) {
        NewLevel(1, setGame);
    }
    if (game.array) {
        game.getTubesList().forEach((tube, index) => {
            renderedTubes.push(<Tube tube={tube} key={index} tubeIndex={index} selected={selected == index} handleTubeSelection={handleTubeSelection} />);
        });
    }

    return (
        <div className="container">
            <h1>Level {game.id}</h1>
            {renderedTubes}
        </div>
    )
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
