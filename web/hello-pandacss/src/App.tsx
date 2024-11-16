import { css } from "../styled-system/css";
import { Box } from "../styled-system/jsx";

function App() {
  return (
    <>
      <div className={css({ fontSize: "2xl", fontWeight: "bold" })}>
        Hello 🐼!
      </div>

      {/**
       * Box
       *
       * ボックスパターンには追加のスタイルは含まれない。
       * function形式ではcss関数と同等。これはJSX形式で使用すると便利で、
       * styled.divコンポーネントと同等で、主にJSXで使用可能なスタイルの
       * 小道具を取得するために機能する。
       *
       * @see https://panda-css.com/docs/concepts/patterns#box
       */}
      <h2>Box</h2>

      <Box color="blue.300">
        <div>Cool!</div>
      </Box>
    </>
  );
}

export default App;
