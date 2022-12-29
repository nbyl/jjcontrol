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
        await updateRoomState()
    }, 3000);

    async function updateRoomState() {
        try {
            const response = await fetch("/api/room");
            const data = JSON.parse(await response.text());

            setRoomName(data.name)
            setLightOn(data.lightOn);
        } catch (err) {
            console.log(err)
        }
    }

    async function switchLight(event: React.MouseEvent<HTMLDivElement>, on: boolean) {
        try {
            await fetch("/api/room", {
                method: 'PUT',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({
                    'lightOn': on,
                })
            });
            setLightOn(on)
        } catch (err) {
            console.log(err)
        }
    }

    return (
        <div className="App">
            <h1>{roomName}</h1>

            {lightOn && (
                <div style={{color: 'yellow', cursor: 'pointer'}} onClick={(e) => switchLight(e, false)}>
                    <Icon size={256} icon={ic_lightbulb}/>
                </div>
            )}
            {!lightOn && (
                <div style={{cursor: 'pointer'}} onClick={(e) => switchLight(e, true)}>
                    <Icon size={256} icon={ic_lightbulb_outline_twotone}/>
                </div>
            )}
        </div>
    );
}

export default App
