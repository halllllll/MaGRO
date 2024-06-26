import { type StyleFunctionProps, extendTheme } from '@chakra-ui/react';

const theme = extendTheme({
  // Chakra UIのコンポーネントは基本的にstyles.globalを見ず、themeのfontsセクションに定義されたフォントを参照する。
  fonts: {
    body: '"BIZ UDPGothic", Meiryo, sans-serif',
    heading: '"BIZ UDPGothic", Meiryo, sans-serif',
    mono: '"BIZ UDPGothic", Meiryo, sans-serif',
  },
  styles: {
    global: {
      'html, body': {
        fontFamily: '"BIZ UDPGothic",Meiryo, sans-serif',
      },
      '::-webkit-scrollbar': {
        width: '0.5rem',
        bgColor: 'transparent',
      },
      '::-webkit-scrollbar-thumb': {
        bgColor: 'rgba(183, 194, 218, 0.12)',
      },
      '*': {
        scrollbarWidth: 'thin',
        scrollbarColor: 'rgba(183, 194, 218, 0.12) transparent',
        scrollbarGutter: 'stable',
      },
    },
  },
  components: {
    Heading: {
      baseStyle: {
        padding: { base: '.5rem 0', md: '1rem 0' },
        marginBottom: '1rem',
        textAlign: 'center',
      },
      defaultProps: {
        as: 'h2',
        size: 'lg',
      },
    },
    Link: {
      baseStyle: (props: StyleFunctionProps) => ({
        color: props.colorMode === 'dark' ? 'whiteAlpha.900' : '#17242A',
        _hover: {
          color: props.colorMode === 'dark' ? '#b7c2da' : '#314e89',
          textDecoration: 'none',
        },
      }),
    },
    Button: {
      defaultProps: {
        colorScheme: 'facebook',
      },
    },
    Card: {
      defaultProps: {
        variant: 'outline',
      },
      baseStyle: {
        container: {
          padding: '1.2rem',
          gap: '1.2rem',
        },
      },
    },
  },
});

export { theme };
