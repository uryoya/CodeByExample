import { css } from "../styled-system/css";
import { Button } from "./components/ui/button";

function App() {
  return (
    <>
      <div className={css({ fontSize: "2xl", fontWeight: "bold" })}>
        Hello 🐼!
      </div>
      <Button>Park UI Button !</Button>
    </>
  );
}

export default App;
