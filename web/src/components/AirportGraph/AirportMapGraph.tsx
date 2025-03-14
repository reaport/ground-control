import React, { FC, useMemo } from "react"
// reactflow для визуализации
import ReactFlow, {
  Node as FlowNode,
  Edge as FlowEdge,
  useNodesState,
  useEdgesState,
  Position,
} from "reactflow"
import "reactflow/dist/style.css"

// Компоненты из shadcn (пример с Card)
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card"
import { AirportMap } from "@/api"

interface AirportMapGraphProps {
  data: AirportMap
}

export const AirportMapGraph: FC<AirportMapGraphProps> = ({ data }) => {
  /**
   * Простой пример размещения: раскладываем все узлы
   * равномерно по окружности вокруг некоего "центра"
   */
  const centerX = 500
  const centerY = 400
  const radius = 300

  // Превращаем ваши Node[] в формат React Flow
  const nodes: FlowNode[] = useMemo(() => {
    const total = data.nodes.length
    return data.nodes.map((node, index) => {
      const angle = (2 * Math.PI * index) / total
      const x = centerX + radius * Math.cos(angle)
      const y = centerY + radius * Math.sin(angle)

      return {
        id: node.id,
        position: { x, y },
        // Текст внутри узла – можно сделать красивее
        data: { label: node.id },
        // Небольшие стили, чтобы узлы были похожи на «карточки»
        style: {
          backgroundColor: "white",
          border: "1px solid #ccc",
          borderRadius: 6,
          padding: 8,
          width: 140,
          textAlign: "center",
        },
      }
    })
  }, [data])

  // Превращаем ваши Edge[] в формат React Flow
  const edges: FlowEdge[] = useMemo(() => {
    return data.edges.map((edge) => ({
      id: `${edge.from}-${edge.to}`,
      source: edge.from,
      target: edge.to,
      // Можно подписать длину как label
      label: `${edge.distance}m`,
      // Добавляем «стрелочку» на конце
      markerEnd: { type: "arrowclosed" },
      style: { strokeWidth: 1.5 },
      labelBgPadding: [6, 3],
      labelBgBorderRadius: 4,
      labelBgStyle: { fill: "rgba(255,255,255,0.8)" },
    }))
  }, [data])

  // React Flow позволяет отслеживать изменения узлов/рёбер
  const [nodesState, , onNodesChange] = useNodesState(nodes)
  const [edgesState, , onEdgesChange] = useEdgesState(edges)

  return (
    <Card className="w-full h-[80vh]">
      <CardHeader>
        <CardTitle>Airport Graph</CardTitle>
      </CardHeader>
      <CardContent className="relative w-full h-full">
        <ReactFlow
          nodes={nodesState}
          edges={edgesState}
          onNodesChange={onNodesChange}
          onEdgesChange={onEdgesChange}
          fitView
          fitViewOptions={{ padding: 0.2 }}
          // Можно включить перетаскивание (pan) и зум колёсиком
          panOnScroll
          zoomOnScroll
        />
      </CardContent>
    </Card>
  )
}
