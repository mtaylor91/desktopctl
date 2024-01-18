import axios from 'axios';
import { render } from 'preact';
import { useEffect, useState } from 'preact/hooks';
import './style.css';


function VM({ vm: {
  metadata: { name },
  spec: { running },
  status: { printableStatus }
}, refresh, token }) {
  const bgClass = running ? 'bg-green' : 'bg-red';
  return (
    <div class={`vm ${bgClass}`}>
      <div class="vm-name">{name}</div>
      <div class="vm-status">{printableStatus}</div>
      <div class="vm-actions">
        <VMActions name={name} running={running} refresh={refresh} token={token} />
      </div>
    </div>
  );
}


function VMActions({ name, running, refresh, token }) {
  const startVM = evt => {
    evt.preventDefault();
    console.log(`Starting ${name}`);
    axios.post(`/api/vms/${name}/start`, undefined, { headers: {
      'Authorization': `Bearer ${token}`
    }}).then(({ data }) => {
      refresh();
    }).catch(err => {
      console.error(err);
      if (err.response.status === 401) {
        localStorage.removeItem('token');
        // Redirect to login page
        window.location.href = '/auth/login';
      }
    })
  }

  const stopVM = evt => {
    evt.preventDefault();
    console.log(`Stopping ${name}`);
    axios.post(`/api/vms/${name}/stop`, undefined, { headers: {
      'Authorization': `Bearer ${token}`
    }}).then(({ data }) => {
      refresh();
    }).catch(err => {
      console.error(err);
      if (err.response.status === 401) {
        localStorage.removeItem('token');
        // Redirect to login page
        window.location.href = '/auth/login';
      }
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


function App({ token }) {
  const [vms, setVMs] = useState([]);

  const listVMs = () => {
    axios.get('/api/vms', { headers: {
      'Authorization': `Bearer ${token}`
    }}).then(({ data }) => {
      setVMs(data.items);
    }).catch(err => {
      console.error(err);
      if (err.response.status === 401) {
        localStorage.removeItem('token');
        // Redirect to login page
        window.location.href = '/auth/login';
      }
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
        {vms.map((vm, i) => <VM key={i} refresh={listVMs} token={token} vm={vm} />)}
      </div>
    </>
  );
}


(async () => {
	// Check for oauth2 token in local storage or url
	const urlParams = new URLSearchParams(window.location.search);
	const token = urlParams.get('token') || localStorage.getItem('token');

	if (!token) {
		// Redirect to login page
		window.location.href = '/auth/login';
	} else {
		// Set token in local storage
		localStorage.setItem('token', token);
		// Render app
		render(<App token={token} />, document.getElementById('app'));
	}
})();
