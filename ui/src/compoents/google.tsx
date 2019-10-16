import * as React from "react";

export const GoogleLink: React.FunctionComponent = () => {
  const [message, setMessage] = React.useState("");

  React.useEffect(() => {
    window.renderer.on("testing", (evt, val) => {
      console.log(`Recieve event "${evt}" from golang`);
      setMessage(JSON.stringify(val));
    });
  });

  return (
    <div>
      <a href="http://google.com">Google</a>
      <p>
        Event: <b>testing</b>
      </p>
      <p>Message from go: {message}</p>
    </div>
  );
};
