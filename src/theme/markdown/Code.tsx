import React from 'react';
import { HighlightedCode } from '@mui/docs/HighlightedCode';
import { styled, alpha, useTheme } from '@mui/material/styles';
import { BrandingProvider } from '@mui/docs/branding';
import { SxProps, Button } from '@mui/material';
import clipboardCopy from 'clipboard-copy';
import Box from '@mui/material/Box';
import Stack from '@mui/material/Stack';

interface CodeProps {
    codeString: string;
    sx?: SxProps; // 添加 sx 属性，并设置为可选
}

export default function Code({ codeString, sx }: CodeProps) {
    const theme = useTheme();
    const { copy, isCopied } = useClipboardCopy();
    const key = 'Ctrl + ';
    return (
        <>
            <Stack direction={{ xs: 'column', }}
                spacing={1}
                useFlexGap
                sx={{ pt: 1, width: { xs: '100%' }, margin: "center", }}>
                <HighlightedCode code={codeString} language="jsx" copyButtonHidden={true} sx={{
                    margin: "center", textAlign: "center",
                    color: theme.palette.text.secondary,
                    ...theme.applyStyles('dark', {
                        color: '#1A2027',
                    }),
                }} />
            </Stack>
            <Stack direction={{ xs: 'column', sm: 'row', }}
                useFlexGap
                sx={{ pt: 1, width: { xs: '100%' }, justifyContent: 'center' }}>
                <Button
                    variant="contained"
                    color='success'
                    onClick={async () => {
                        await copy(codeString);
                    }}
                >
                    {isCopied ? 'Copied' : 'Copy'}
                </Button>
            </Stack>
        </>
    );
}

function useClipboardCopy() {
    const [isCopied, setIsCopied] = React.useState(false);
    const timeout = React.useRef<ReturnType<typeof setTimeout> | undefined>(undefined);

    React.useEffect(
        () => () => {
            clearTimeout(timeout.current);
        },
        [],
    );

    const copy = async (text: string) => {
        await clipboardCopy(text);
        setIsCopied(true);
        clearTimeout(timeout.current);
        timeout.current = setTimeout(() => {
            setIsCopied(false);
        }, 1200);
    };

    return { copy, isCopied };
}
