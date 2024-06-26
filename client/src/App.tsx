import { useState } from 'react';

type Res = {
  answer: string;
  forced: boolean;
};

type Ping = {
  status: number;
  timestamp: string;
  message: string;
};

function App() {
  const [apiRes, setApiRes] = useState<Res>();
  const [data, setData] = useState<Ping>();

  const externalApiBtnHandler = async () => {
    fetch('/api1')
      .then((res) => {
        if (!res.ok) {
          throw new Error('Network response was not ok');
        }
        return res.json();
      })
      .then((data) => {
        console.info(data);
        setApiRes(data);
      })
      .catch((err) => {
        console.error(err);
      });
  };

  const localApiBtnHandler = async () => {
    const res = await fetch('/api/hello');
    if (!res.ok) {
      console.error(res.status, ' -- ', res.statusText);
    } else {
      const data = (await res.json()) as Ping;
      console.dir(data);
      setData(data);
    }
  };

  return (
    <div>
      <div>
        <button type={'button'} onClick={externalApiBtnHandler}>
          API1
        </button>
        <div>{apiRes?.answer}</div>
      </div>
      <button type={'button'} onClick={localApiBtnHandler}>
        API2
      </button>
      {data && (
        <div>
          result: {data.message}, ${data.timestamp}
        </div>
      )}
    </div>
  );
}

export default App;
