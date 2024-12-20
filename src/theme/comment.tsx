import React, { useEffect, useRef } from 'react';
import {
    type WalineInstance,
    init,
} from '@waline/client';
import { useTheme } from '@mui/material/styles'; // 导入 useTheme
import Container from '@mui/material/Container';
import { Box } from '@mui/material';

import '@waline/client/style';

const WalineComment = () => {
    const walineInstanceRef = useRef<WalineInstance | null>(null);
    const theme = useTheme();

    useEffect(() => {
        // 初始化 Waline
        walineInstanceRef.current = init({
            el: '#waline',
            pageview: true, 
            serverURL: 'https://comment.gitproxy.click',
        });

        return () => walineInstanceRef.current?.destroy();
    }, []);

    useEffect(() => {
        if (walineInstanceRef.current) {
            walineInstanceRef.current.update({ dark: theme.palette.mode === 'dark' });
        }
        document.documentElement.style.setProperty('--waline-bg-color', theme.palette.background.default);
        
        if (theme.palette.mode === 'dark') {
            document.documentElement.style.setProperty('--waline-bg-color', theme.palette.background.paper);
        }
    }, [theme]); 

    return (
        <Container>
 
            <Box
                id="waline"
                sx={{
                    width: '100%',
                    flexDirection: { xs: 'column', sm: 'row' },
                    justifyContent: 'space-between',
                    bgcolor: 'transparent', 
                    p: 2,  
                    borderRadius: 2,  
                    boxShadow: theme.shadows[2],  
                }}
            />  
        </Container>
    );
};

export default WalineComment;
