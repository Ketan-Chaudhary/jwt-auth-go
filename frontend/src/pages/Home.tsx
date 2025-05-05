import { useEffect, useState } from "react";

const Home = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  useEffect(() => {
    (async () => {
      const response = await fetch("http://localhost:8000/api/user", {
        method: "GET",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });

      const content = await response.json();
      setName(content.name);
      setEmail(content.email);
    })();
  });

  const submit = async () => {
    await fetch("http://localhost:8000/api/logout", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });

    setName("");
    setEmail("");
  };
  return (
    <div>
      <p> {name ? "Hi" + name : "you are not log in"} </p>
      <p> {email ? "login with" + email : ""}</p>

      {name && email ? (
        <button onClick={submit} type="submit">
          Logout
        </button>
      ) : (
        ""
      )}
    </div>
  );
};

export default Home;
