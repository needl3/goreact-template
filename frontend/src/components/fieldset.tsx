export default function CustomFieldSet({
  legendText,
  children,
  legendClass,
  fieldsetClass
}: {
  legendText: string;
  children: React.ReactNode;
  legendClass?: string;
  fieldsetClass?: string
}) {
  return (
    <fieldset className={`relative w-full ${fieldsetClass}`}>
      <legend
        className={`absolute -top-2 left-2 text-sm bg-white px-1 font-semibold capitalize ${legendClass}`}
      >
        {legendText}
      </legend>
      {children}
    </fieldset>
  );
}
