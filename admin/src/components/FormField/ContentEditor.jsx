import { Form } from "antd";
import { Editor } from '@tinymce/tinymce-react';
import { tokenStorage } from '@/utils/user-storage';
import { request } from '@umijs/max';
import qs from 'qs';

// function virtualInput(){
//   const input = document.createElement('input');
//   input.setAttribute('type', 'file');
//   input.setAttribute('accept', 'image/*');

//   input.addEventListener('change', (e) => {
//     const file = e.target.files[0];

//     const reader = new FileReader();
//     reader.addEventListener('load', () => {
//       /*
//         Note: Now we need to register the blob in TinyMCEs image blob
//         registry. In the next release this part hopefully won't be
//         necessary, as we are looking to handle it internally.
//       */
//       const id = 'blobid' + (new Date()).getTime();
//       const blobCache =  tinymce.activeEditor.editorUpload.blobCache;
//       const base64 = reader.result.split(',')[1];
//       const blobInfo = blobCache.create(id, file, base64);
//       blobCache.add(blobInfo);

//       /* call the callback and populate the Title field with the file name */
//       cb(blobInfo.blobUri(), { title: file.name });
//     });
//     reader.readAsDataURL(file);
//   });

//   input.click();
// }

async function postUploadFormImage(formData) {
  // const formData = new FormData()
  // formData.append('file',file)
  // formData.append('path',path)
  const res = await request('/api/admin/upload/form-file', {
      method: 'POST',
      data:formData,
  });
  return res.data
}

function chooseFile(accept="image/*"){
  return new Promise(function(resolve, reject){
      var input = document.createElement('input')
      input.accept=accept
      input.type="file"
      // input.onchange = function(event){
      //     resolve(input.files[0])
      //     input.value = ''
      // }
      input.addEventListener('change', async (e) => {
        resolve(e.target.files[0])
      })
      input.click()
  })
}

const TinyMCEApiKey = 'eay9qssigezvotd1b504f031zaa1dgql4ejczjrcce4gm6zo'

export default function ContentEditor({ value, onChange,path="editor",privacy="public",height=640  }){

    // https://ant-design.antgroup.com/components/form-cn#formitemusestatus
    //const { status, errors } = Form.Item.useStatus();
    //const [value,setValue] = useState('<p>hello world</p>')
    //const [text, setText] = useState('');

    //const token = tokenStorage.get()
    //const headers = { "Authorization": "Bearer " + token }
    //const apiUrl = API_URL + "/api/admin/upload/form-image"

    const images_upload_handler =  async (blobInfo, progress) => {
      // https://www.tiny.cloud/docs/tinymce/6/upload-images/#images_upload_handler
      // https://www.tiny.cloud/docs/tinymce/6/php-upload-handler/
      // console.log(blobInfo,progress)

      const formData = new FormData();
      formData.append('file', blobInfo.blob(), blobInfo.filename());
      formData.append('path', path);
      formData.append('privacy', privacy);
      const res = await postUploadFormImage(formData)
      return res.url
    }

    const file_picker_callback = function(callback,value,meta){
      //console.log(value,meta)
      // todo: 模拟input选择文件
      // callback('https://www.baidu.com/img/bd_logo1.png', { text: 'My text' });

      let accept = "*"
      if(meta.filetype == 'file'){
        accept = "*"
      }else if(meta.filetype == 'image'){
        accept = "image/*"
      }else if(meta.filetype == 'media'){
        accept = "audio/*,video/*"
      }
      //console.log(accept)
      const input = document.createElement('input');
      input.setAttribute('type', 'file');
      input.setAttribute('accept', accept);
      input.addEventListener('change', async (e) => {
        const file = e.target.files[0];

        //console.log(file)
        const formData = new FormData();
        formData.append('file', file,file.name);
        formData.append('path', path);
        formData.append('privacy', privacy);
        const res = await postUploadFormImage(formData)
        callback(res.url,{"title":file.name})
      })
      input.click()
    }

    return (
      <Editor
        apiKey={TinyMCEApiKey}
        tinymceScriptSrc='/tinymce/tinymce.min.js'
        value={value}
        onInit={(evt, editor) => {
          // setText(editor.getContent({format: 'text'}));
        }}
        onEditorChange={(newValue, editor) => {
          // setValue(newValue);
          // setText(editor.getContent({format: 'text'}));
          onChange(newValue)
        }}
        init={{
          //placeholder: '请输入...',
          placeholder: '',

          images_upload_handler: images_upload_handler,
          //images_upload_base_path: "http://127.0.0.1",
          //images_upload_url: apiUrl,

          // https://www.tiny.cloud/docs/tinymce/6/file-image-upload/#adding-a-file-browser-to-image-and-media-dialogs
          file_picker_callback: file_picker_callback,

          // https://www.tiny.cloud/docs/tinymce/6/ui-localization/#language
          language_url: '/tinymce/langs/zh-Hans.js',
          language: "zh-Hans",

          // https://www.tiny.cloud/docs/tinymce/6/full-featured-open-source-demo/
          menubar: false,
          plugins: 'preview importcss searchreplace autolink autosave save directionality code visualblocks visualchars fullscreen image link media template codesample table charmap pagebreak nonbreaking anchor insertdatetime advlist lists wordcount help charmap quickbars emoticons accordion',
          // https://www.tiny.cloud/docs/tinymce/6/available-toolbar-buttons/#the-core-toolbar-buttons
          // toolbar: "undo redo | accordion accordionremove | blocks fontfamily fontsize | bold italic underline strikethrough | align numlist bullist | link image | table media | lineheight outdent indent| forecolor backcolor removeformat | charmap emoticons | code fullscreen preview | save print | pagebreak anchor codesample | ltr rtl",
          toolbar: "undo redo | blocks fontfamily fontsize | bold italic underline strikethrough | align numlist bullist | link image | table media | lineheight outdent indent| forecolor backcolor removeformat | charmap emoticons | code fullscreen preview | print | pagebreak anchor codesample | ltr rtl",

          // 关闭自定义右键菜单
          contextmenu: false,

          // 编辑器高度
          height: height,

          // 行开头点击弹出快捷菜单
          quickbars_insert_toolbar:"",

          // // 预设链接、图片链接、模板HTML
          // link_list: [
          //   { title: '网站首页', value: 'https://www.baidu.com' },
          //   { title: 'My page 2', value: 'http://www.moxiecode.com' }
          // ],
          // image_list: [
          //   { title: 'LOGO图片', value: 'https://www.baidu.com/img/bd_logo1.png' },
          //   { title: 'My page 2', value: 'http://www.moxiecode.com' }
          // ],
          // templates: [
          //   { title: 'New Table', description: 'creates a new table', content: '<div class="mceTmpl"><table width="98%%"  border="0" cellspacing="0" cellpadding="0"><tr><th scope="col"> </th><th scope="col"> </th></tr><tr><td> </td><td> </td></tr></table></div>' },
          //   { title: 'Starting my story', description: 'A cure for writers block', content: 'Once upon a time...' },
          //   { title: 'New list with dates', description: 'New List with dates', content: '<div class="mceTmpl"><span class="cdate">cdate</span><br><span class="mdate">mdate</span><h2>My List</h2><ul><li></li><li></li></ul></div>' }
          // ],

          // // 自定义按钮
          // // customInsertButton customDateButton template 
          // // https://www.tiny.cloud/docs/tinymce/6/custom-basic-toolbar-button/
          // setup: (editor) => {
          //   editor.ui.registry.addButton('customInsertButton', {
          //     text: 'My Button',
          //     onAction: (_) => editor.insertContent(`&nbsp;<strong>It's my button!</strong>&nbsp;`)
          //   });
        
          //   const toTimeHtml = (date) => `<time datetime="${date.toString()}">${date.toDateString()}</time>`;
          //   editor.ui.registry.addButton('customDateButton', {
          //     icon: 'insert-time',
          //     tooltip: 'Insert Current Date',
          //     enabled: false,
          //     onAction: (_) => editor.insertContent(toTimeHtml(new Date())),
          //     onSetup: (buttonApi) => {
          //       const editorEventCallback = (eventApi) => {
          //         buttonApi.setEnabled(eventApi.element.nodeName.toLowerCase() !== 'time');
          //       };
          //       editor.on('NodeChange', editorEventCallback);
        
          //       /* onSetup should always return the unbind handlers */
          //       return () => editor.off('NodeChange', editorEventCallback);
          //     }
          //   });
          // },
        }}
      />
      );
}