import {useState} from 'react'
import './App.css'
import {useInterval} from "./hooks/useInterval";
import {Icon} from 'react-icons-kit'
import {ic_lightbulb} from 'react-icons-kit/md/ic_lightbulb'
import {ic_lightbulb_outline_twotone} from 'react-icons-kit/md/ic_lightbulb_outline_twotone'

function App() {
    const [roomName, setRoomName] = useState("");
    const [lightOn, setLightOn] = useState(false);

    useInterval(async () => {
        const fetchData = async () => {
            const response = await fetch("/api/room");
            const data = JSON.parse(await response.text());

            setRoomName(data.name)
            setLightOn(data.lightOn);
        };

        fetchData().catch((err) => console.log(err));
    }, 3000);

    return (
        <div className="App">
            <h1>{roomName}</h1>

            {lightOn && (
                <div style={{color: 'yellow'}}>
                    <Icon size={256} icon={ic_lightbulb}/>
                </div>
            )}
            {!lightOn && (
                <div>
                    <Icon size={256} icon={ic_lightbulb_outline_twotone}/>
                </div>
            )}
        </div>
    );
}

export default App
