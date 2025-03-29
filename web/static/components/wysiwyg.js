const boldIconPath = new URL("./icons/bold.svg", import.meta.url).href;
const italicIconPath = new URL("./icons/italic.svg", import.meta.url).href;
const underlinedIconPath = new URL("./icons/underlined.svg", import.meta.url)
  .href;
const stricThroughIconPath = new URL(
  "./icons/stric_through.svg",
  import.meta.url,
).href;
const alignLeftIconPath = new URL("./icons/align_left.svg", import.meta.url)
  .href;
const alignRightIconPath = new URL("./icons/align_right.svg", import.meta.url)
  .href;
const alignCenterIconPath = new URL("./icons/align_center.svg", import.meta.url)
  .href;
const alignJustifyIconPath = new URL(
  "./icons/align_justify.svg",
  import.meta.url,
).href;
const orderListIconPath = new URL("./icons/ordered_list.svg", import.meta.url);
const unorderedListIconPath = new URL(
  "./icons/unorderd_list.svg",
  import.meta.url,
);

const outdentIconPath = new URL("./icons/outdent.svg", import.meta.url).href;
const indentIconPath = new URL("./icons/indent.svg", import.meta.url).href;

const template = document.createElement("template");
template.innerHTML = `
<style>
  .container {
    height: 100%;
    display: flex;
    flex-direction: column;
    padding: 2px;
  }

  .toolbar {
    display: flex;
    flex-direction: row;
    position: relative;
  }

  .button {
    padding: 2px;
    background-color: inherit;
    border: 0px;

    &:hover {
      filter: brightness(0.8);
      cursor: pointer;
    }

  }

  .icon {
    margin: auto auto;
  }

  .menu {
    display: flex;
    flex-direction: column;
    z-index: 1;
    background-color: inherit;
    position: relative;

    .submenu {
      margin:   0px;
      display: none;
    }

    &:hover {
      .submenu {
        display: block;
        position: absolute;
        top: 100%; 
        left: 0; 
        z-index: 10;
        background-color: inherit;
        border: 1px solid black;
      }
    }
  }
</style>
<div class="container">
  <slot  name="label"></slot>
  <div class="toolbar" part="toolbar">
    <button part="button" class="button" data-action="bold" data-tag-name="b" title="Bold">
      <img height="full" width="full" src="${boldIconPath}" alt="B"/>
    </button>
    <button part="button" class="button" data-action="italic" data-tag-name="i" title="Italic">
      <img class="icon" src="${italicIconPath}" alt="I"/>
    </button>
    <button  part="button" class="button" data-action="underline" data-tag-name="u" title="Underline">
      <img src="${underlinedIconPath}" alt="U"/>
    </button>
    <button  part="button" class="button" data-action="strikeThrough" data-tag-name="strike" title="Strike through">
      <img src="${stricThroughIconPath}" alt="S"/>
    </button>

    <div class="menu">
      <button class="button" part="button" data-action="justifyLeft" data-style="textAlign:left" title="Justify left">
        <img part="icon" src="${alignLeftIconPath}" alt="Left"/>
      </button>
      <div class="submenu">
        <button class="button" part="button" data-action="justifyCenter" data-style="textAlign:center" title="Justify center">
          <img src="${alignCenterIconPath}" alt="Center"/>
        </button>
        <button class="button" part="button" data-action="justifyRight" data-style="textAlign:right" title="Justify right">
          <img src="${alignRightIconPath}" alt="Right"/>
        </button>
        <button class="button" part="button" data-action="justifyFull" data-style="textAlign:justify" title="Justify full">
          <img src="${alignJustifyIconPath}" alt="Full"/>
        </button>
      </div>
    </div>

    <button part="button" class="button" data-action="insertOrderedList" data-tag-name="ol" title="Ordered list">
      <img src="${orderListIconPath}" alt="OL"/>
    </button>
    <button part="button" class="button" data-action="insertUnorderedList" data-tag-name="ul" title="Unordered list">
      <img src="${unorderedListIconPath}" alt="UL"/>
    </button>
    <button part="button" class="button" data-action="outdent"  title="Outdent">
      <img src="${outdentIconPath}" alt="outdent"/>
    </button>
    <button part="button" class="button" data-action="indent"  title="Outdent">
      <img src="${indentIconPath}" alt="indent"/>
    </button>
  </div>
  <slot id="wysiwyg-editor-message"class="message" name="message"></slot>
</div>
`;

class WYSIWYGEditor extends HTMLElement {
  constructor() {
    super();
    const shadow = this.attachShadow({ mode: "open" });
    shadow.appendChild(template.content.cloneNode(true));
  }
}

customElements.define("wysiwyg-editor", WYSIWYGEditor);
