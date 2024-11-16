import { css } from "../styled-system/css";
import { Box, Container } from "../styled-system/jsx";
import { container } from "../styled-system/patterns";

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

      {/**
       * Container
       *
       * コンテナパターンは最大幅のコンテナを作成し、コンテンツを中央に配置する
       * ために使用される。
       * デフォルトでは、コンテナは次のスタイルが適用される。
       *
       * maxWidth: 8xl
       * marginX: auto
       * position: relative
       * paddingX: { base: 4, md: 6, lg: 8 }
       *
       * @see https://panda-css.com/docs/concepts/patterns#container
       */}
      <h2>Container</h2>

      <h3>function style</h3>
      <div className={container()}>
        <div>First</div>
        <div>Second</div>
        <div>Third</div>
      </div>

      <h3>JSX style</h3>
      <Container>
        <div>First</div>
        <div>Second</div>
        <div>Third</div>
      </Container>
    </>
  );
}

export default App;
