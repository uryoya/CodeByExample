import React from "react";
import { css } from "../styled-system/css";
import {
  AspectRatio,
  Box,
  Container,
  Flex,
  HStack,
  Stack,
  VStack,
  Wrap,
} from "../styled-system/jsx";
import {
  aspectRatio,
  container,
  flex,
  hstack,
  stack,
  vstack,
  wrap,
} from "../styled-system/patterns";

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
      <Box color="blue.300" className={css({ border: "solid" })}>
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
      <div className={container({ border: "solid" })}>
        <div className={css({ border: "solid", color: "red" })}>First</div>
        <div className={css({ border: "solid", color: "red" })}>Second</div>
        <div className={css({ border: "solid", color: "red" })}>Third</div>
      </div>

      <h3>JSX style</h3>
      <Container className={css({ border: "solid" })}>
        <div className={css({ border: "solid", color: "red" })}>First</div>
        <div className={css({ border: "solid", color: "red" })}>Second</div>
        <div className={css({ border: "solid", color: "red" })}>Third</div>
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
      <div className={stack({ gap: "6", padding: "4", border: "solid" })}>
        <div className={css({ border: "solid", color: "red" })}>First</div>
        <div className={css({ border: "solid", color: "red" })}>Second</div>
        <div className={css({ border: "solid", color: "red" })}>Third</div>
      </div>

      <h3>JSX style</h3>
      <Stack gap="6" padding="4" className={css({ border: "solid" })}>
        <div className={css({ border: "solid", color: "red" })}>First</div>
        <div className={css({ border: "solid", color: "red" })}>Second</div>
        <div className={css({ border: "solid", color: "red" })}>Third</div>
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
      <div className={hstack({ gap: "6", border: "solid" })}>
        <div className={css({ border: "solid", color: "red" })}>First</div>
        <div className={css({ border: "solid", color: "red" })}>Second</div>
        <div className={css({ border: "solid", color: "red" })}>Third</div>
      </div>

      <h3>JSX style</h3>
      <HStack gap="6" className={css({ border: "solid" })}>
        <div className={css({ border: "solid", color: "red" })}>First</div>
        <div className={css({ border: "solid", color: "red" })}>Second</div>
        <div className={css({ border: "solid", color: "red" })}>Third</div>
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
      <div className={vstack({ gap: "6", border: "solid" })}>
        <div className={css({ border: "solid", color: "red" })}>First</div>
        <div className={css({ border: "solid", color: "red" })}>Second</div>
        <div className={css({ border: "solid", color: "red" })}>Third</div>
      </div>

      <h3>JSX style</h3>
      <VStack gap="6" className={css({ border: "solid" })}>
        <div className={css({ border: "solid", color: "red" })}>First</div>
        <div className={css({ border: "solid", color: "red" })}>Second</div>
        <div className={css({ border: "solid", color: "red" })}>Third</div>
      </VStack>

      {/**
       * Wrap
       *
       * Wrap(折り返し)パターンは要素間にスペースを追加するために使用され、十分なスペースがない場合は自動的に折り返す。
       * Wrapパターンは次のプロパティを受け付ける。
       * - gap: The gap between the elements in the stack.
       * - columnGap: The gap between the elements in the stack horizontally.
       * - rowGap: The gap between the elements in the stack vertically.
       * - align: An alias for the css align-items property.
       * - justify: An alias for the css justify-content property.
       *
       * @see https://panda-css.com/docs/concepts/patterns#wrap
       */}
      <SectionTitle>Wrap</SectionTitle>

      <h3>function style</h3>
      <div className={wrap({ gap: "6", border: "solid" })}>
        <div className={css({ width: "300px", border: "solid", color: "red" })}>
          First
        </div>
        <div className={css({ width: "300px", border: "solid", color: "red" })}>
          Second
        </div>
        <div className={css({ width: "300px", border: "solid", color: "red" })}>
          Third
        </div>
      </div>

      <h3>JSX style</h3>
      <Wrap gap="6" className={css({ border: "solid" })}>
        <div style={{ width: "300px", border: "solid", color: "red" }}>
          First
        </div>
        <div style={{ width: "300px", border: "solid", color: "red" }}>
          Second
        </div>
        <div style={{ width: "300px", border: "solid", color: "red" }}>
          Third
        </div>
      </Wrap>

      {/**
       * Aspect Ratio
       *
       * Aspect Ratio (アスペクト比) パターンは、固定アスペクト比のコンテナを作成するために使用される。
       * 画像、地図、ビデオなどのメディアを表示するときに使用される。
       *
       * @note ほとんどの場合、パターンの代わりに、aspectRatio プロパティを使用することが推奨される。
       * @see https://panda-css.com/docs/concepts/patterns#aspect-ratio
       */}
      <SectionTitle>Aspect Ratio</SectionTitle>

      <h3>function style</h3>
      <div className={aspectRatio({ ratio: 16 / 9, border: "solid" })}>
        <iframe
          src="https://www.google.com/maps/embed"
          title="Google map"
          frameBorder="0"
        />
      </div>

      <h3>JSX style</h3>
      <AspectRatio ratio={16 / 9} className={css({ border: "solid" })}>
        <iframe
          src="https://www.google.com/maps/embed"
          title="Google map"
          frameBorder="0"
        />
      </AspectRatio>

      {/**
       * Flex
       *
       * FlexパターンはFlexコンテナーの作成に使用され、Flexプロパティのいくつかのショートカットを提供する。
       * Flexパターンは次のプロパティを受け付ける。
       * - direction: The flex direction of the container. Can be row, column, row-reverse or column-reverse.
       * - wrap: Whether to wrap the flex items. The value is a boolean.
       * - align: An alias for the css align-items property.
       * - justify: An alias for the css justify-content property.
       * - basis: An alias for the css flex-basis property.
       * - grow: An alias for the css flex-grow property.
       * - shrink: An alias for the css flex-shrink property.
       *
       * @see https://panda-css.com/docs/concepts/patterns#flex
       */}
      <SectionTitle>Flex</SectionTitle>

      <h3>function style</h3>
      <div
        className={flex({
          direction: "row",
          align: "center",
          border: "solid",
        })}
      >
        <div className={css({ border: "solid", color: "red" })}>First</div>
        <div className={css({ border: "solid", color: "red" })}>Second</div>
        <div className={css({ border: "solid", color: "red" })}>Third</div>
      </div>

      <h3>JSX style</h3>
      <Flex direction="row" align="center" className={css({ border: "solid" })}>
        <div className={css({ border: "solid", color: "red" })}>First</div>
        <div className={css({ border: "solid", color: "red" })}>Second</div>
        <div className={css({ border: "solid", color: "red" })}>Third</div>
      </Flex>
    </>
  );
}

export default App;
