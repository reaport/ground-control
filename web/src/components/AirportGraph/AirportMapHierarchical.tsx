import React from "react"
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card"
// Установите: npm install dagre
import * as dagre from "dagre"
import { AirportMap } from "@/api"

type Props = {
  airportMap: AirportMap
}

export function AirportMapHierarchical({ airportMap }: Props) {
  const { nodes, edges } = airportMap

  // Создаем граф dagre
  const g = new dagre.graphlib.Graph()
    .setGraph({
      // можно 'LR' = слева-направо, 'TB' = сверху-вниз и т.п.
      rankdir: "TB",
      // настройте отступы для более "просторного" вида
      ranksep: 100, 
      edgesep: 50,  
      nodesep: 50, 
    })
    .setDefaultEdgeLabel(() => ({}))

  // Добавляем узлы
  nodes.forEach((node) => {
    // Размер «прямоугольника» для dagre. Он нужен, чтобы dagre знал, 
    // какой "бокс" выделять для узла. Ширину/высоту можно указать на глаз.
    g.setNode(node.id, { width: 100, height: 40 })
  })

  // Добавляем рёбра
  edges.forEach((edge) => {
    g.setEdge(edge.from, edge.to)
  })

  // Вызываем авто-раскладку
  dagre.layout(g)

  // g.node(u) теперь содержит { x, y } для каждой вершины u
  // Соберём в Map для удобства
  const layoutedNodes = new Map<string, { x: number; y: number }>()
  g.nodes().forEach((nodeId) => {
    const nodeWithPos = g.node(nodeId) as dagre.Node
    layoutedNodes.set(nodeId, { x: nodeWithPos.x, y: nodeWithPos.y })
  })

  // Найдем границы, чтобы "подтянуть" диаграмму к верхнему левому углу
  const allX = Array.from(layoutedNodes.values()).map((p) => p.x)
  const allY = Array.from(layoutedNodes.values()).map((p) => p.y)
  const minX = Math.min(...allX)
  const minY = Math.min(...allY)
  const maxX = Math.max(...allX)
  const maxY = Math.max(...allY)

  // Сделаем небольшой padding вокруг
  const padding = 50
  const width = maxX - minX + padding * 2
  const height = maxY - minY + padding * 2

  return (
    <Card>
      <CardHeader>
        <CardTitle>Иерархическая раскладка AirportMap</CardTitle>
      </CardHeader>
      <CardContent>
        <svg
          width={width}
          height={height}
          style={{ border: "1px solid #ccc", background: "#fff" }}
        >
          {/* Рисуем рёбра */}
          {edges.map((edge, index) => {
            const fromPos = layoutedNodes.get(edge.from)!
            const toPos = layoutedNodes.get(edge.to)!
            return (
              <line
                key={index}
                x1={fromPos.x - minX + padding}
                y1={fromPos.y - minY + padding}
                x2={toPos.x - minX + padding}
                y2={toPos.y - minY + padding}
                stroke="#999"
                strokeWidth={2}
              />
            )
          })}

          {/* Рисуем узлы */}
          {nodes.map((node) => {
            const pos = layoutedNodes.get(node.id)!
            const cx = pos.x - minX + padding
            const cy = pos.y - minY + padding

            return (
              <g key={node.id}>
                {/* Прямоугольник (или круг) – на ваш вкус */}
                <rect
                  x={cx - 50} // 100 ширина / 2
                  y={cy - 20} // 40 высота / 2
                  width={100}
                  height={40}
                  rx={6}
                  ry={6}
                  fill="#fff"
                  stroke="#333"
                  strokeWidth={1}
                  style={{ cursor: "pointer" }}
                />
                {/* Текст по центру */}
                <text
                  x={cx}
                  y={cy + 5}
                  fill="#333"
                  fontSize={12}
                  textAnchor="middle"
                  dominantBaseline="middle"
                  style={{ pointerEvents: "none", userSelect: "none" }}
                >
                  {node.id}
                </text>
              </g>
            )
          })}
        </svg>
      </CardContent>
    </Card>
  )
}
