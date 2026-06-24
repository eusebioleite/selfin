with prods as (
	select
        pro_codpro,
        tpp_descri,
        gru_descri,
        pro_unimed,
        pro_descri
	from f_prods
	join f_cadgru
	    on gru_codigo = pro_grprod
	join f_tpprod
	    on tpp_codigo = pro_tpprod
	where exists (select 1 from f_prods where pro_descri like '%RESINA%')
),
estoq as (
	select
        est_codpro,
        est_locest,
        est_sldant,
        est_qtentr,
        est_qtcons,
        est_qtcomp
	from f_estoq
	where exists (select 1 from f_prods where pro_descri like '%RESINA%')
),
lotemat as (
	select
        lmt_codpro,
        lmt_locest,
        lmt_qtdlot,
        lmt_sailot,
        lmt_qtddev,
        lmt_lotfor,
        lmt_observ,
        lmt_tiplot,
        lmt_corrid,
        lmt_datval,
        lmt_datfab
	from f_lotemat
	where exists (select 1 from f_prods where pro_descri like '%RESINA%')
),
kardex as (
	select
        kar_codpro,
        kar_locest,
        sum(kar_qtdmov) as total_consumo_of
	from f_kardex
	where kar_tipdoc in ( 'OF', 'OFC' )
	and exists (select 1 from f_prods where pro_descri like '%RESINA%')
	group by kar_codpro, kar_locest
)
select
pro_codpro,
pro_descri,
tpp_descri,
gru_descri,
pro_unimed,
nvl(est_locest, 'SEM LOCAL') as locest,
nvl((est_sldant + est_qtentr - est_qtcons), 0) as saldo_logico_atual,
nvl(est_qtcomp, 0) as quantidade_comprometida,
nvl(((est_sldant + est_qtentr - est_qtcons) - est_qtcomp), 0) as saldo_disponivel,
lmt_lotfor as lote_fornecedor,
nvl((lmt_qtdlot - lmt_sailot + lmt_qtddev), 0) as saldo_fisico_lote,
lmt_datfab,
lmt_datval,
case
    when lmt_datval is null then 'Lote sem data de validade'
    when lmt_datval < trunc(sysdate) then 'VENCIDO'
    when ( lmt_datval - trunc(sysdate) ) <= 30 then '0-30 Dias (Crítico)'
    when ( lmt_datval - trunc(sysdate) ) <= 60 then '31-60 Dias (Alerta)'
    else '61+ Dias (Estável)'
end as vida,
ceil(lmt_datval - trunc(sysdate)) as dias_para_vencer,
nvl(total_consumo_of, 0) as consumo_total_ordem_fab
from prods
left join estoq
    on est_codpro = pro_codpro
left join lotemat
    on lmt_codpro = est_codpro and lmt_locest = est_locest
left join kardex
on kar_codpro = est_codpro and kar_locest = est_locest
order by saldo_logico_atual asc, pro_codpro asc