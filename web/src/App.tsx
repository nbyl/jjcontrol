import {useState} from 'react'
import './App.css'
import {useInterval} from "./hooks/useInterval";

function App() {
    const [data, setData] = useState("");

    useInterval(async () => {
        const fetchData = async () => {
            const response = await fetch("/api/room");
            const data = await response.text();
            setData(data);
        };

        fetchData().catch((err) => console.log(err));
    }, 3000);

    return (
        <div className="App">
            <h1>{data}</h1>
        </div>
    );
}

export default App
