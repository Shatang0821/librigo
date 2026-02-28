function App() {
  return (
    <div className="min-h-screen bg-neutral-950 flex flex-col items-center justify-center p-4">
      {/* カード部分 */}
      <div className="max-w-md w-full bg-neutral-900 border border-neutral-800 rounded-2xl p-8 shadow-2xl transition-transform hover:scale-[1.02]">
        
        {/* ステータスバッジ */}
        <div className="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-emerald-500/10 border border-emerald-500/20 mb-6">
          <div className="w-2 h-2 rounded-full bg-emerald-500 animate-pulse" />
          <span className="text-xs font-medium text-emerald-500 uppercase tracking-wider">Environment Ready</span>
        </div>

        {/* メインテキスト */}
        <h1 className="text-3xl font-bold text-white mb-2">
          Tailwind CSS v4 <span className="text-blue-500">Active</span>
        </h1>
        <p className="text-neutral-400 mb-8">
          Node v22 + React 19 + Tailwind v4 の最新環境が正常に動作しています。
        </p>

        {/* テスト用グリッド */}
        <div className="grid grid-cols-3 gap-2">
          <div className="h-12 bg-blue-500 rounded-lg shadow-lg shadow-blue-500/20" />
          <div className="h-12 bg-emerald-500 rounded-lg shadow-lg shadow-emerald-500/20" />
          <div className="h-12 bg-purple-500 rounded-lg shadow-lg shadow-purple-500/20" />
        </div>
      </div>

      <p className="mt-8 text-neutral-600 text-sm italic">
        デザインが崩れていなければ、導入成功です！
      </p>
    </div>
  )
}

export default App