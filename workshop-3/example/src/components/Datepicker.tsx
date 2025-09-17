import { useState, useMemo, useRef, useEffect } from 'react';

interface CalendarDay {
  date: Date;
  otherMonth: boolean;
}

const Datepicker = () => {
  const [selectedDate, setSelectedDate] = useState<Date | null>(null);
  const [showCalendar, setShowCalendar] = useState(false);
  const [currentMonth, setCurrentMonth] = useState(new Date().getMonth());
  const [currentYear, setCurrentYear] = useState(new Date().getFullYear());
  const datepickerRef = useRef<HTMLDivElement>(null);

  const formattedDate = useMemo(() => 
    selectedDate ? selectedDate.toLocaleDateString() : '', 
    [selectedDate]
  );

  const weekdays = ['Su', 'Mo', 'Tu', 'We', 'Th', 'Fr', 'Sa'];

  const daysInMonth = (year: number, month: number) =>
    new Date(year, month + 1, 0).getDate();

  const calendarDays = useMemo(() => {
    const days: CalendarDay[] = [];
    const firstDayOfWeek = new Date(currentYear, currentMonth, 1).getDay();
    
    // Previous month days
    const prevMonthDays = daysInMonth(currentYear, currentMonth - 1);
    for (let i = firstDayOfWeek - 1; i >= 0; i--) {
      days.push({
        date: new Date(currentYear, currentMonth - 1, prevMonthDays - i),
        otherMonth: true,
      });
    }
    
    // Current month days
    const thisMonthDays = daysInMonth(currentYear, currentMonth);
    for (let i = 1; i <= thisMonthDays; i++) {
      days.push({ 
        date: new Date(currentYear, currentMonth, i), 
        otherMonth: false 
      });
    }
    
    // Next month days (fill to 6 weeks grid)
    const nextDays = 42 - days.length;
    for (let i = 1; i <= nextDays; i++) {
      days.push({ 
        date: new Date(currentYear, currentMonth + 1, i), 
        otherMonth: true 
      });
    }
    
    return days;
  }, [currentYear, currentMonth]);

  const currentMonthName = useMemo(() =>
    new Date(currentYear, currentMonth).toLocaleString('default', { month: 'long' }),
    [currentYear, currentMonth]
  );

  const toggleCalendar = () => setShowCalendar(!showCalendar);

  const prevMonth = () => {
    if (currentMonth === 0) {
      setCurrentMonth(11);
      setCurrentYear(currentYear - 1);
    } else {
      setCurrentMonth(currentMonth - 1);
    }
  };

  const nextMonth = () => {
    if (currentMonth === 11) {
      setCurrentMonth(0);
      setCurrentYear(currentYear + 1);
    } else {
      setCurrentMonth(currentMonth + 1);
    }
  };

  const selectDate = (day: CalendarDay) => {
    if (day.otherMonth) return;
    setSelectedDate(day.date);
    setShowCalendar(false);
  };

  const isSelected = (day: CalendarDay) =>
    selectedDate &&
    day.date.toDateString() === selectedDate.toDateString();

  // Close calendar when clicking outside
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (datepickerRef.current && !datepickerRef.current.contains(event.target as Node)) {
        setShowCalendar(false);
      }
    };

    if (showCalendar) {
      document.addEventListener('mousedown', handleClickOutside);
    }

    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, [showCalendar]);

  return (
    <div className="relative inline-block" ref={datepickerRef}>
      <input
        type="text"
        readOnly
        value={formattedDate}
        onClick={toggleCalendar}
        placeholder="Select date"
        className="w-40 p-2 border border-neutral-300 rounded-md cursor-pointer 
                   focus:outline-none focus:ring-2 focus:ring-web-green-500 
                   bg-white text-neutral-700 placeholder-neutral-400"
      />
      
      {showCalendar && (
        <div className="absolute top-full left-0 mt-2 bg-white border border-neutral-300 
                        rounded-lg shadow-lg z-50 p-4 min-w-64">
          {/* Header */}
          <div className="flex justify-between items-center mb-4">
            <button 
              onClick={prevMonth} 
              className="p-2 hover:bg-neutral-100 rounded-md text-neutral-600 
                         hover:text-neutral-800 transition-colors"
            >
              &#8249;
            </button>
            <span className="font-semibold text-neutral-800 text-lg">
              {currentMonthName} {currentYear}
            </span>
            <button 
              onClick={nextMonth} 
              className="p-2 hover:bg-neutral-100 rounded-md text-neutral-600 
                         hover:text-neutral-800 transition-colors"
            >
              &#8250;
            </button>
          </div>
          
          {/* Weekdays */}
          <div className="grid grid-cols-7 text-center mb-2">
            {weekdays.map((day) => (
              <span 
                key={day} 
                className="font-semibold text-neutral-600 py-2 text-sm"
              >
                {day}
              </span>
            ))}
          </div>
          
          {/* Days */}
          <div className="grid grid-cols-7 text-center gap-1">
            {calendarDays.map((day, index) => (
              <button
                key={`${day.date.getTime()}-${index}`}
                onClick={() => selectDate(day)}
                type="button"
                className={`
                  p-2 cursor-pointer rounded-full text-sm transition-colors
                  ${day.otherMonth 
                    ? 'text-neutral-300 hover:text-neutral-400' 
                    : 'text-neutral-700 hover:bg-web-green-100'
                  }
                  ${isSelected(day) 
                    ? 'bg-web-green-500 text-white hover:bg-web-green-600' 
                    : ''
                  }
                `}
              >
                {day.date.getDate()}
              </button>
            ))}
          </div>
        </div>
      )}
    </div>
  );
};

export default Datepicker;
