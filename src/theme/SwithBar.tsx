import * as React from 'react';
import { useState } from 'react';
import { useTheme } from '@mui/material/styles'; // 导入 useTheme
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField'; // 导入 TextField
import Code from './markdown/Code';
import { Stack } from '@mui/material';
import Alert from '@mui/material/Alert';

interface CodeSpaceProps {
    codeString: string;
    index: number;
    value: number;
    inputValue: string;
    notice: string;
}

const CodeSpace: React.FC<CodeSpaceProps> = ({ codeString, index, value, inputValue, notice }) => {
    function getCopyCode(index ){
        if (index==0){
            return "git clone https://gitproxy.click/"+codeString
        }
        if (index==1){
            return "wget https://gitproxy.click/"+codeString  +"\n"+"curl https://gitproxy.click/" +codeString
        }
        if (index==2){
            return "https://gitproxy.click/"+codeString
        }
        return ""
    }
    return (
        <>
            <Box hidden={value !== index} width={"100%"}    >
                <Alert sx={{ justifyContent: 'center' }}  color='success'>{notice} </Alert>

            </Box>
            <Box hidden={value !== index || inputValue == ""} width={"100%"} margin={"center"}>
                <Code codeString={getCopyCode(index)} sx={{ }} />
            </Box>
        </>
    );
};
interface SwitchBarProps {
    inputValue: string;
}

export default function SwitchBar({ inputValue }: SwitchBarProps) {
    const [value, setValue] = useState(0);

    const handleChange = (event: React.SyntheticEvent, newValue: number) => {
        setValue(newValue);
    };

    return (
        <>
            <Tabs
               color="primary"
        // size="small"
                value={value}
                onChange={handleChange}
                aria-label="basic tabs example"
                sx={{ display: 'flex', justifyContent: 'space-between', width: '100%' }}
            >
                <Tab   color="primary" label="Git Clone" sx={{ flexGrow: 1, textAlign: 'center' }} />
                <Tab  color="primary"  label="Wget && Curl" sx={{ flexGrow: 1, textAlign: 'center' }} />
                <Tab   color="primary" label="Direct Download" sx={{ flexGrow: 1, textAlign: 'center' }} />
            </Tabs>


            <CodeSpace value={value} index={0}
                codeString={`${inputValue}`} inputValue={inputValue} notice='
                支持直接 git clone 获取仓库  请输入仓库地址  复制粘贴命令 ' />
            <CodeSpace value={value} index={1} codeString={`${inputValue}`} inputValue={inputValue} notice='
             支持直接根据文件地址获取下载链接 请输入仓库地址 复制粘贴命令 ' />
            <CodeSpace value={value} index={2} codeString={`${inputValue}`} inputValue={inputValue} notice='
            支持从浏览器直接下载 输入下载文件地址 点击开始下载即可下载' />
        </>
    );
}