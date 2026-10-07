package view

const (
	snippetMermaidHead = `---
title: "{{title}}"
---
%% This diagram has been generated with go-structurizr
%% https://github.com/krzysztofreczek/go-structurizr
%% title: {{title}}
flowchart TB
`
	snippetMermaidClassDef = `  classDef {{shape_style}} fill:{{background_color_hash}},color:{{font_color_hash}},stroke:{{border_color_hash}}
`
	snippetMermaidComponent = `
  {{node}}
  class {{component_id}} {{shape_style}}
`
	snippetMermaidConnection = `  {{component_id_from}} -.-> {{component_id_to}}
`
	snippetMermaidLinkStyle = `  linkStyle default stroke:{{line_color_hash}}
`
	paramNode = "{{node}}"
)
