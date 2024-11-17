import React from "react";
import { css } from "../styled-system/css";
import { Box, Container, HStack, Stack, VStack } from "../styled-system/jsx";
import { container, hstack, stack, vstack } from "../styled-system/patterns";

type SectionTitleProps = {
  children?: React.ReactNode;
};

const SectionTitle: React.FC<SectionTitleProps> = ({ children }) => (
  <h2 className={css({ bg: "gray.200", fontSize: "xl" })}>{children}</h2>
);

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
      <SectionTitle>Box</SectionTitle>
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
      <SectionTitle>Container</SectionTitle>

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

      {/**
       * Stack
       *
       * スタックパターンは、水平または垂直のスタックを作成するために使用される。
       * Stackは次のプロパティを受け付ける。
       * - direction: An alias for the css flex-direction property. Default is column.
       * - gap: The gap between the elements in the stack.
       * - align: An alias for the css align-items property.
       * - justify: An alias for the css justify-content property.
       *
       * @see https://panda-css.com/docs/concepts/patterns#stack
       */}
      <SectionTitle>Stack</SectionTitle>

      <h3>function style</h3>
      <div className={stack({ gap: "6", padding: "4" })}>
        <div>First</div>
        <div>Second</div>
        <div>Third</div>
      </div>

      <h3>JSX style</h3>
      <Stack gap="6" padding="4">
        <div>First</div>
        <div>Second</div>
        <div>Third</div>
      </Stack>

      {/**
       * HStack
       *
       * HStackパターンはStackパターンのラッパーで、directionプロパティがhorizontalに設定され、要素を垂直方向に中央揃えにする。
       *
       * @see https://panda-css.com/docs/concepts/patterns#hstack
       */}
      <SectionTitle>HStack</SectionTitle>

      <h3>function style</h3>
      <div className={hstack({ gap: "6" })}>
        <div>First</div>
        <div>Second</div>
        <div>Third</div>
      </div>

      <h3>JSX style</h3>
      <HStack gap="6">
        <div>First</div>
        <div>Second</div>
        <div>Third</div>
      </HStack>

      {/**
       * VStack
       *
       * VStackパターンはStackパターンのラッパーで、directionプロパティがverticalに設定され、要素を水平方向に中央揃えにする。
       *
       * @see https://panda-css.com/docs/concepts/patterns#vstack
       */}
      <SectionTitle>VStack</SectionTitle>

      <h3>function style</h3>
      <div className={vstack({ gap: "6" })}>
        <div>First</div>
        <div>Second</div>
        <div>Third</div>
      </div>

      <h3>JSX style</h3>
      <VStack gap="6">
        <div>First</div>
        <div>Second</div>
        <div>Third</div>
      </VStack>
    </>
  );
}

export default App;
