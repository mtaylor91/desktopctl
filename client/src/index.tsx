import axios from 'axios';
import { render } from 'preact';
import { useEffect, useState } from 'preact/hooks';
import './style.css';

function VM({ vm: {
  metadata: { name },
  spec: { running },
  status: { printableStatus }
}, refresh }) {
  const bgClass = running ? 'bg-green' : 'bg-red';
  return (
    <div class={`vm ${bgClass}`}>
      <div class="vm-name">{name}</div>
      <div class="vm-status">{printableStatus}</div>
      <div class="vm-actions">
        <VMActions name={name} running={running} refresh={refresh} />
      </div>
    </div>
  );
}

function VMActions({ name, running, refresh }) {
  const startVM = evt => {
    evt.preventDefault();
    console.log(`Starting ${name}`);
    axios.post(`/api/vms/${name}/start`).then(({ data }) => {
      refresh();
    }).catch(err => {
      console.error(err);
    })
  }

  const stopVM = evt => {
    evt.preventDefault();
    console.log(`Stopping ${name}`);
    axios.post(`/api/vms/${name}/stop`).then(({ data }) => {
      refresh();
    }).catch(err => {
      console.error(err);
    })
  }

  if (running) {
    return (
      <button onClick={stopVM}>
        Stop
      </button>
    );
  } else {
    return (
      <button onClick={startVM}>
        Start
      </button>
    );
  }
}

export function App() {
  const [vms, setVMs] = useState([]);

  const listVMs = () => {
    axios.get('/api/vms').then(({ data }) => {
      setVMs(data.items);
    });
  };

  useEffect(() => {
    const interval = setInterval(() => listVMs(), 5000);
    listVMs();
    return () => clearInterval(interval);
  }, []);

  return (
    <>
      <h1>Virtual Machines</h1>
      <div class="vms">
        {vms.map((vm, i) => <VM key={i} refresh={listVMs} vm={vm}/>)}
      </div>
    </>
  );
}

render(<App />, document.getElementById('app'));
