import Datepicker from './components/Datepicker';

function App() {

  return (
    <>
      {/* Header Section */}
      <header className="bg-web-green-600 text-white py-6 px-8">
        <div className="max-w-6xl mx-auto">
          <h1 className="text-3xl font-bold">Design Token Example</h1>
          <p className="text-web-green-100 mt-2">แสดงตัวอย่างการใช้ design token จาก Tailwind CSS Theme</p>
        </div>
      </header>

      {/* Main Content */}
      <main className="bg-neutral-50 min-h-screen py-12 px-8">
        <div className="max-w-6xl mx-auto">
          
          {/* Datepicker Demo Section */}
          <section className="mb-16">
            <h2 className="text-2xl font-bold text-neutral-900 mb-8">Datepicker Component</h2>
            <div className="bg-white p-6 rounded-lg shadow">
              <div className="mb-4">
                <div className="block text-sm font-medium text-neutral-700 mb-2">
                  เลือกวันที่:
                </div>
                <Datepicker />
              </div>
              <p className="text-neutral-600 text-sm">
                ตัวอย่าง Datepicker component ที่แปลงมาจาก Vue และใช้ design token ของโปรเจค
              </p>
            </div>
          </section>
          
          {/* Color Palette Section */}
          <section className="mb-16">
            <h2 className="text-2xl font-bold text-neutral-900 mb-8">Color Palette</h2>
            
            {/* Brand Colors */}
            <div className="mb-8">
              <h3 className="text-lg font-semibold text-neutral-700 mb-4">Brand Colors</h3>
              <div className="grid grid-cols-2 md:grid-cols-5 gap-4">
                <div className="bg-web-green-100 p-4 rounded-md text-center">
                  <div className="text-sm font-medium text-neutral-700">Green 100</div>
                </div>
                <div className="bg-web-green-300 p-4 rounded-md text-center">
                  <div className="text-sm font-medium text-neutral-700">Green 300</div>
                </div>
                <div className="bg-web-green-500 p-4 rounded-md text-center text-white">
                  <div className="text-sm font-medium">Green 500</div>
                </div>
                <div className="bg-web-green-600 p-4 rounded-md text-center text-white">
                  <div className="text-sm font-medium">Green 600</div>
                </div>
                <div className="bg-web-green-800 p-4 rounded-md text-center text-white">
                  <div className="text-sm font-medium">Green 800</div>
                </div>
              </div>
            </div>

            {/* Functional Colors */}
            <div className="mb-8">
              <h3 className="text-lg font-semibold text-neutral-700 mb-4">Functional Colors</h3>
              <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                <div className="bg-success p-4 rounded-md text-center text-white">
                  <div className="text-sm font-medium">Success</div>
                </div>
                <div className="bg-warning p-4 rounded-md text-center text-white">
                  <div className="text-sm font-medium">Warning</div>
                </div>
                <div className="bg-error p-4 rounded-md text-center text-white">
                  <div className="text-sm font-medium">Error</div>
                </div>
                <div className="bg-info p-4 rounded-md text-center text-white">
                  <div className="text-sm font-medium">Info</div>
                </div>
              </div>
            </div>
          </section>

          {/* Typography Section */}
          <section className="mb-16">
            <h2 className="text-2xl font-bold text-neutral-900 mb-8">Typography</h2>
            <div className="space-y-4 bg-white p-6 rounded-lg shadow">
              <div className="text-5xl font-bold text-neutral-900">Heading 1 (5xl)</div>
              <div className="text-4xl font-bold text-neutral-700">Heading 2 (4xl)</div>
              <div className="text-3xl font-bold text-neutral-700">Heading 3 (3xl)</div>
              <div className="text-2xl font-semibold text-neutral-700">Heading 4 (2xl)</div>
              <div className="text-xl font-semibold text-neutral-700">Heading 5 (xl)</div>
              <div className="text-base text-neutral-700">Body Text (base) - นี่คือข้อความปกติที่ใช้ในเนื้อหาหลัก</div>
              <div className="text-sm text-neutral-500">Small Text (sm) - ข้อความขนาดเล็กสำหรับรายละเอียดเพิ่มเติม</div>
              <div className="text-xs text-neutral-400">Extra Small (xs) - ข้อความขนาดพิเศษเล็ก</div>
            </div>
          </section>

          {/* Components Section */}
          <section className="mb-16">
            <h2 className="text-2xl font-bold text-neutral-900 mb-8">UI Components</h2>
            
            {/* Buttons */}
            <div className="mb-8">
              <h3 className="text-lg font-semibold text-neutral-700 mb-4">Buttons</h3>
              <div className="flex flex-wrap gap-4">
                <button className="bg-web-green-500 hover:bg-web-green-600 text-white px-6 py-3 rounded-md font-medium transition-colors">
                  Primary Button
                </button>
                <button className="bg-neutral-200 hover:bg-neutral-300 text-neutral-700 px-6 py-3 rounded-md font-medium transition-colors">
                  Secondary Button
                </button>
                <button className="border border-web-green-500 text-web-green-500 hover:bg-web-green-50 px-6 py-3 rounded-md font-medium transition-colors">
                  Outline Button
                </button>
                <button className="bg-error hover:bg-red-600 text-white px-6 py-3 rounded-md font-medium transition-colors">
                  Danger Button
                </button>
              </div>
            </div>

            {/* Cards */}
            <div className="mb-8">
              <h3 className="text-lg font-semibold text-neutral-700 mb-4">Cards</h3>
              <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                <div className="bg-white p-6 rounded-lg shadow">
                  <h4 className="text-xl font-semibold text-neutral-900 mb-3">Basic Card</h4>
                  <p className="text-neutral-600 mb-4">การ์ดพื้นฐานที่ใช้ shadow และ rounded corners ตาม design token</p>
                  <button className="bg-web-green-500 text-white px-4 py-2 rounded-md text-sm font-medium">
                    Learn More
                  </button>
                </div>
                
                <div className="bg-white p-6 rounded-lg shadow-md border-l-4 border-web-green-500">
                  <h4 className="text-xl font-semibold text-neutral-900 mb-3">Highlighted Card</h4>
                  <p className="text-neutral-600 mb-4">การ์ดที่มีเส้นขอบสีเขียวเพื่อเน้นความสำคัญ</p>
                  <button className="bg-web-green-500 text-white px-4 py-2 rounded-md text-sm font-medium">
                    Learn More
                  </button>
                </div>
                
                <div className="bg-web-green-50 p-6 rounded-lg border border-web-green-200">
                  <h4 className="text-xl font-semibold text-web-green-800 mb-3">Branded Card</h4>
                  <p className="text-web-green-700 mb-4">การ์ดที่ใช้สีจากแบรนด์เป็นพื้นหลัง</p>
                  <button className="bg-web-green-600 text-white px-4 py-2 rounded-md text-sm font-medium">
                    Learn More
                  </button>
                </div>
              </div>
            </div>

            {/* Alerts */}
            <div>
              <h3 className="text-lg font-semibold text-neutral-700 mb-4">Alerts</h3>
              <div className="space-y-4">
                <div className="bg-green-50 border border-success text-green-800 p-4 rounded-md">
                  <strong>Success:</strong> การดำเนินการเสร็จสิ้นเรียบร้อยแล้ว
                </div>
                <div className="bg-yellow-50 border border-warning text-yellow-800 p-4 rounded-md">
                  <strong>Warning:</strong> กรุณาตรวจสอบข้อมูลอีกครั้ง
                </div>
                <div className="bg-red-50 border border-error text-red-800 p-4 rounded-md">
                  <strong>Error:</strong> เกิดข้อผิดพลาดในการประมวลผล
                </div>
                <div className="bg-blue-50 border border-info text-blue-800 p-4 rounded-md">
                  <strong>Info:</strong> ข้อมูลเพิ่มเติมสำหรับผู้ใช้งาน
                </div>
              </div>
            </div>
          </section>

          {/* Spacing Examples */}
          <section className="mb-16">
            <h2 className="text-2xl font-bold text-neutral-900 mb-8">Spacing Examples</h2>
            <div className="bg-white p-6 rounded-lg shadow">
              <div className="space-y-4">
                <div className="p-1 bg-web-green-100 rounded text-center text-sm">Spacing 1 (4px)</div>
                <div className="p-2 bg-web-green-100 rounded text-center text-sm">Spacing 2 (8px)</div>
                <div className="p-4 bg-web-green-100 rounded text-center text-sm">Spacing 4 (16px)</div>
                <div className="p-6 bg-web-green-100 rounded text-center text-sm">Spacing 6 (24px)</div>
                <div className="p-8 bg-web-green-100 rounded text-center text-sm">Spacing 8 (32px)</div>
              </div>
            </div>
          </section>

        </div>
      </main>

      {/* Footer */}
      <footer className="bg-neutral-800 text-white py-8 px-8">
        <div className="max-w-6xl mx-auto text-center">
          <p className="text-neutral-300">
            ตัวอย่างการใช้ Design Token ใน Tailwind CSS Theme Configuration
          </p>
          <p className="text-neutral-400 text-sm mt-2">
            สร้างด้วย React + TypeScript + Tailwind CSS
          </p>
        </div>
      </footer>
    </>
  )
}

export default App
